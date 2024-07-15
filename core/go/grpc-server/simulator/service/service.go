package simulator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	au "github.com/hijjiri/simulator/core/go/grpc-server/aura"
	b "github.com/hijjiri/simulator/core/go/grpc-server/battle"
	d "github.com/hijjiri/simulator/core/go/grpc-server/deck"
	e "github.com/hijjiri/simulator/core/go/grpc-server/extension"
	h "github.com/hijjiri/simulator/core/go/grpc-server/hero"
	lc "github.com/hijjiri/simulator/core/go/grpc-server/landcryptid"
	pb "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	"github.com/hijjiri/simulator/core/go/grpc-server/user"
)

const (
	workerNum = 10
	queueSize = 1
)

var requestQueue chan func()

func init() {
	requestQueue = make(chan func(), queueSize)
	for i := 0; i < workerNum; i++ {
		go worker()
	}
}

func worker() {
	for req := range requestQueue {
		req()
	}
}

type Server struct {
	pb.UnimplementedSimulatorServiceServer
}

func (s *Server) SimulateBattle(ctx context.Context, in *pb.SimulateRequest) (*pb.SimulateResponse, error) {
	log.Printf("Received SimulateBattle request: %v", in)

	resultCh := make(chan *pb.SimulateResponse, 1)
	errorCh := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	requestQueue <- func() {
		defer wg.Done()
		log.Println("Processing request in worker")

		// ユーザーIDをランダムに生成
		rand.Seed(time.Now().UnixNano())
		userID := rand.Uint32()

		userData, err := user.LockUser(ctx, userID)
		if err != nil {
			errorCh <- err
			log.Printf("Error locking user: %v", err)
			return
		}

		log.Printf("User ID: %d, Is Locked: %v", userData.Uid, userData.IsLocked)

		if !userData.IsLocked {
			err := errors.New("user is not properly locked")
			errorCh <- err
			log.Println(err)
			return
		}

		log.Printf("Simulating battle %d times", in.GetCounts())
		log.Printf("Offense Deck: %v", in.GetSimulateOffenseDeck().GetUnits())
		log.Printf("Defense Deck: %v", in.GetSimulateDefenseDeck().GetUnits())

		// シミュレーションの結果を生成
		rand.Seed(time.Now().UnixNano())
		battleID := rand.Uint32()
		attacker := rand.Uint32()
		defender := rand.Uint32()
		resultCounts := make(map[uint32]int32)

		// NewBattleUnitFromDeck
		offenseUnits, err := createBattleUnits(ctx, convertDeckUnits(in.GetSimulateOffenseDeck().GetUnits()), true)
		if err != nil {
			errorCh <- err
			log.Printf("Error creating offense units: %v", err)
			return
		}
		log.Println("Created offense units")

		defenseUnits, err := createBattleUnits(ctx, convertDeckUnits(in.GetSimulateDefenseDeck().GetUnits()), false)
		if err != nil {
			errorCh <- err
			log.Printf("Error creating defense units: %v", err)
			return
		}
		log.Println("Created defense units")

		units := append(offenseUnits, defenseUnits...)
		if len(units) < 6 {
			err := fmt.Errorf("units number not enough %v", len(units))
			errorCh <- err
			log.Println(err)
			return
		}
		log.Println("Units are sufficient")

		for i := 0; i < int(in.GetCounts()); i++ {
			winner := rand.Uint32()
			resultCounts[winner]++
		}
		log.Println("Battle simulation completed")

		// ロック解除
		_, unlockErr := user.UnlockUser(ctx, userID)
		if unlockErr != nil {
			log.Printf("Failed to unlock user: %v", unlockErr)
		} else {
			log.Printf("User %d unlocked successfully", userID)
		}

		resultCh <- &pb.SimulateResponse{
			BattleId:     battleID,
			Attacker:     attacker,
			Defender:     defender,
			ResultCounts: resultCounts,
		}
		log.Println("SimulateBattle response sent")
	}

	// ゴルーチンでワーカーグループの完了を待つ
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
		log.Println("Worker group completed")
	}()

	select {
	case res := <-resultCh:
		log.Println("Returning SimulateBattle response")
		return res, nil
	case err := <-errorCh:
		log.Printf("Error occurred: %v", err)
		return nil, err
	case <-ctx.Done():
		log.Println("Context done")
		return nil, ctx.Err()
	}
}

func createBattleUnits(ctx context.Context, deckUnits []*d.DeckUnit, isAttacker bool) ([]*b.Unit, error) {
	units := make([]*b.Unit, len(deckUnits))

	for i, deckUnit := range deckUnits {
		position := int32(i)
		heroData := createHeroData(deckUnit.HeroId)
		extension1 := createExtensionData(deckUnit.ExtensionIds[0])
		extension2 := createExtensionData(deckUnit.ExtensionIds[1])
		cryptidInfo := &lc.CryptidInfo{}
		auras := []*au.AuraEffect{}
		contentType := "simulator"

		log.Printf("Creating unit at position %d with heroData: %v, extension1: %v, extension2: %v", position, heroData, extension1, extension2)
		units[i] = b.NewBattleUnitFromDeck(position, deckUnit, heroData, extension1, extension2, cryptidInfo, auras, deckUnit.SkillOrders, deckUnit.HeroActiveIndex, contentType)
		log.Printf("Created unit at position %d: %v", position, units[i])
	}

	return units, nil
}

func convertDeckUnits(pbUnits []*pb.DeckUnit) []*d.DeckUnit {
	deckUnits := make([]*d.DeckUnit, len(pbUnits))
	for i, pbUnit := range pbUnits {
		deckUnits[i] = &d.DeckUnit{
			HeroId:          pbUnit.HeroId,
			ExtensionIds:    pbUnit.ExtensionIds,
			SkillOrders:     pbUnit.SkillOrders,
			HeroActiveIndex: pbUnit.HeroActiveIndex,
		}
		log.Printf("Converted DeckUnit %d: %v", i, deckUnits[i])
	}
	return deckUnits
}

func createHeroData(heroId uint32) *h.HeroData {
	heroData := &h.HeroData{
		HeroId: heroId,
		Param: &h.Param{
			Hp:  1000,
			Phy: 100,
			Int: 100,
			Agi: 100,
		},
		StockedSkillIds: []uint32{1, 2, 3},
		Passive:         1,
	}
	log.Printf("Created HeroData: %v", heroData)
	return heroData
}

func createExtensionData(extensionId uint32) *e.ExtensionData {
	extensionData := &e.ExtensionData{
		ExtensionId: extensionId,
		Param: &e.Param{
			Hp:  100,
			Phy: 10,
			Int: 10,
			Agi: 10,
		},
		Active: 1,
	}
	log.Printf("Created ExtensionData: %v", extensionData)
	return extensionData
}
