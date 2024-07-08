package simulator

import (
	"context"
	pb "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	"log"
	"math/rand"
	"time"
)

type Server struct {
	pb.UnimplementedSimulatorServiceServer
}

func (s *Server) SimulateBattle(ctx context.Context, in *pb.SimulateRequest) (*pb.SimulateResponse, error) {
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

	return &pb.SimulateResponse{
		BattleId:     battleID,
		Attacker:     attacker,
		Defender:     defender,
		ResultCounts: resultCounts,
	}, nil
}
