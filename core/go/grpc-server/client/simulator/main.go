package main

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
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
	client := pb.NewSimulatorServiceClient(conn)

	var wg sync.WaitGroup
	requests := 1 // 100回リクエストを送信

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			_, err := client.SimulateBattle(ctx, &pb.SimulateRequest{
				Counts: 1,
				SimulateOffenseDeck: &pb.Deck{
					Units: []*pb.DeckUnit{
						{HeroId: 1, ExtensionIds: []uint32{101, 102}, SkillOrders: []int32{0, 1, 2}, HeroActiveIndex: 1},
						{HeroId: 2, ExtensionIds: []uint32{201, 202}, SkillOrders: []int32{0, 2, 1}, HeroActiveIndex: 1},
						{HeroId: 3, ExtensionIds: []uint32{301, 302}, SkillOrders: []int32{1, 2, 0}, HeroActiveIndex: 1},
					},
				},
				SimulateDefenseDeck: &pb.Deck{
					Units: []*pb.DeckUnit{
						{HeroId: 4, ExtensionIds: []uint32{301, 302}, SkillOrders: []int32{0, 1, 2}, HeroActiveIndex: 1},
						{HeroId: 5, ExtensionIds: []uint32{401, 402}, SkillOrders: []int32{0, 2, 1}, HeroActiveIndex: 1},
						{HeroId: 6, ExtensionIds: []uint32{401, 402}, SkillOrders: []int32{1, 2, 0}, HeroActiveIndex: 1},
					},
				},
			})
			if err != nil {
				log.Printf("Request %d failed: %v", i, err)
			} else {
				log.Printf("Request %d succeeded", i)
			}
		}(i)
	}

	wg.Wait()
	log.Println("All requests completed")
}
