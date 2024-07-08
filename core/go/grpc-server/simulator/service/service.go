package simulator

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	"github.com/hijjiri/simulator/core/go/grpc-server/user"
)

const (
	workerNum = 10
	queueSize = 100
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

		rand.Seed(time.Now().UnixNano())
		battleID := rand.Uint32()
		attacker := rand.Uint32()
		defender := rand.Uint32()
		resultCounts := make(map[uint32]int32)

		for i := 0; i < int(in.GetCounts()); i++ {
			winner := rand.Uint32()
			resultCounts[winner]++
		}

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
