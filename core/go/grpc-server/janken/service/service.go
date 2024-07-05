package janken

import (
	"context"
	pb "github.com/hijjiri/simulator/core/go/grpc-server/janken"
	"log"
	"math/rand"
	"time"
)

type Server struct {
	pb.UnimplementedJankenServiceServer
}

func (s *Server) PlayJanken(ctx context.Context, in *pb.JankenRequest) (*pb.JankenResponse, error) {
	log.Printf("Received: %v", in.GetPlayerHand())

	rand.Seed(time.Now().UnixNano())
	serverHand := pb.Hand(rand.Intn(3) + 1)

	result := determineResult(in.GetPlayerHand(), serverHand)

	return &pb.JankenResponse{
		PlayerHand: in.GetPlayerHand(),
		ServerHand: serverHand,
		Result:     result,
	}, nil
}

func determineResult(playerHand pb.Hand, serverHand pb.Hand) string {
	if playerHand == serverHand {
		return "Draw"
	}
	if (playerHand == pb.Hand_ROCK && serverHand == pb.Hand_SCISSORS) ||
		(playerHand == pb.Hand_SCISSORS && serverHand == pb.Hand_PAPER) ||
		(playerHand == pb.Hand_PAPER && serverHand == pb.Hand_ROCK) {
		return "Player Wins"
	}
	return "Server Wins"
}
