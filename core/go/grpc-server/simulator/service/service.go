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
	resultCh := make(chan *pb.SimulateResponse, 1)
	errorCh := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	requestQueue <- func() {
		defer wg.Done()

		// ユーザーIDをランダムに生成
		rand.Seed(time.Now().UnixNano())
		userID := rand.Uint32()

		userData, err := user.LockUser(ctx, userID)
		if err != nil {
			errorCh <- err
			return
		}

		log.Printf("User ID: %d, Is Locked: %v", userData.Uid, userData.IsLocked)

		if !userData.IsLocked {
			errorCh <- errors.New("user is not properly locked")
			return
		}

		log.Printf("Simulating battle %d times", in.GetCounts())
		log.Printf("Offense Deck: %v", in.GetSimulateOffenseDeck().GetUnits())
		log.Printf("Defense Deck: %v", in.GetSimulateDefenseDeck().GetUnits())

		// gRPCから受け取ったDeckUnitデータを変換
		offenseDeckUnits := convertDeckUnits(in.GetSimulateOffenseDeck().GetUnits())
		defenseDeckUnits := convertDeckUnits(in.GetSimulateDefenseDeck().GetUnits())

		// シミュレーションの結果を生成
		rand.Seed(time.Now().UnixNano())
		battleID := rand.Uint32()
		attacker := rand.Uint32()
		defender := rand.Uint32()
		resultCounts := make(map[uint32]int32)

		// NewBattleUnitFromDeck
		offenseUnits, err := createBattleUnits(ctx, offenseDeckUnits, true)
		if err != nil {
			errorCh <- err
			return
		}

		defenseUnits, err := createBattleUnits(ctx, defenseDeckUnits, false)
		if err != nil {
			errorCh <- err
			return
		}

		units := append(offenseUnits, defenseUnits...)
		if len(units) < 6 {
			errorCh <- fmt.Errorf("units number not enough %v", len(units))
			return
		}

		for i := 0; i < int(in.GetCounts()); i++ {
			winner := rand.Uint32()
			resultCounts[winner]++
		}

		// ロック解除
		_, unlockErr := user.UnlockUser(ctx, userID)
		if unlockErr != nil {
			log.Printf("Failed to unlock user: %v", unlockErr)
		}

		resultCh <- &pb.SimulateResponse{
			BattleId:     battleID,
			Attacker:     attacker,
			Defender:     defender,
			ResultCounts: resultCounts,
		}
	}

	// ゴルーチンでワーカーグループの完了を待つ
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()

	select {
	case res := <-resultCh:
		return res, nil
	case err := <-errorCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func createBattleUnits(ctx context.Context, deckUnits []*d.DeckUnit, isAttacker bool) ([]*b.Unit, error) {
	units := make([]*b.Unit, len(deckUnits))

	for i, deckUnit := range deckUnits {
		position := int32(i)
		heroData := createHeroData(deckUnit.HeroId)
		log.Printf("\n\n--- Created HeroData ---\n%+v\n", heroData)

		extension1 := createExtensionData(deckUnit.ExtensionIds[0])
		log.Printf("--- Created ExtensionData ---\n%+v\n", extension1)

		extension2 := createExtensionData(deckUnit.ExtensionIds[1])
		log.Printf("--- Created ExtensionData ---\n%+v\n", extension2)

		cryptidInfo := &lc.CryptidInfo{}
		auras := []*au.AuraEffect{}
		contentType := "simulator"

		log.Printf("\nCreating unit at position %d\nHeroData: %+v\nExtension1: %+v\nExtension2: %+v\n", position, heroData, extension1, extension2)

		units[i] = b.NewBattleUnitFromDeck(position, deckUnit, heroData, extension1, extension2, cryptidInfo, auras, deckUnit.SkillOrders, deckUnit.HeroActiveIndex, contentType)
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
		Aura:   11,
	}
	log.Printf("Created ExtensionData: %v", extensionData)
	return extensionData
}
