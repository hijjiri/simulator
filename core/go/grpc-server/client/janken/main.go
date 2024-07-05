package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/hijjiri/simulator/core/go/grpc-server/janken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewJankenServiceClient(conn)

	playerHand := pb.Hand_ROCK
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "rock":
			playerHand = pb.Hand_ROCK
		case "paper":
			playerHand = pb.Hand_PAPER
		case "scissors":
			playerHand = pb.Hand_SCISSORS
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.PlayJanken(ctx, &pb.JankenRequest{PlayerHand: playerHand})
	if err != nil {
		log.Fatalf("could not play janken: %v", err)
	}

	log.Printf("Player Hand: %v, Server Hand: %v, Result: %v", res.GetPlayerHand(), res.GetServerHand(), res.GetResult())
}
