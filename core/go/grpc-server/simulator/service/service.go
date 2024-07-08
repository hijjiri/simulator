package simulator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	"github.com/hijjiri/simulator/core/go/grpc-server/user"
)

type Server struct {
	pb.UnimplementedSimulatorServiceServer
}

func (s *Server) SimulateBattle(ctx context.Context, in *pb.SimulateRequest) (*pb.SimulateResponse, error) {
	userID := uint32(1001)
	userData, err := user.LockUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	fmt.Println("test", userData.Uid)

	if userData.Uid != 0 {
		return nil, errors.New("user is not properly locked")
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

	for i := 0; i < int(in.GetCounts()); i++ {
		winner := rand.Uint32()
		resultCounts[winner]++
	}

	// ロック解除
	user.UnlockUser(ctx, userID)

	return &pb.SimulateResponse{
		BattleId:     battleID,
		Attacker:     attacker,
		Defender:     defender,
		ResultCounts: resultCounts,
	}, nil
}
