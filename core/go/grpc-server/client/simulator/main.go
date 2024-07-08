package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// デフォルトのシミュレーション回数
	counts := int32(100)

	// コマンドライン引数からシミュレーション回数を取得
	if len(os.Args) > 1 {
		count, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("invalid count value: %v", err)
		}
		counts = int32(count)
	}

	// デッキの作成
	offenseDeck := &pb.Deck{
		Units: []*pb.DeckUnit{
			{HeroId: 1, ExtensionIds: []uint32{101, 102}, SkillOrders: []int32{1, 2}, HeroActiveIndex: 1},
			{HeroId: 2, ExtensionIds: []uint32{201, 202}, SkillOrders: []int32{1, 3}, HeroActiveIndex: 2},
		},
	}
	defenseDeck := &pb.Deck{
		Units: []*pb.DeckUnit{
			{HeroId: 3, ExtensionIds: []uint32{301, 302}, SkillOrders: []int32{2, 3}, HeroActiveIndex: 1},
			{HeroId: 4, ExtensionIds: []uint32{401, 402}, SkillOrders: []int32{1, 2}, HeroActiveIndex: 3},
		},
	}

	// gRPCリクエスト
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SimulateBattle(ctx, &pb.SimulateRequest{
		Counts:              counts,
		SimulateOffenseDeck: offenseDeck,
		SimulateDefenseDeck: defenseDeck,
	})
	if err != nil {
		log.Fatalf("could not simulate battle: %v", err)
	}

	// 結果のフォーマット
	fmt.Printf("Battle ID: %v\n", res.GetBattleId())
	fmt.Printf("Attacker: %v\n", res.GetAttacker())
	fmt.Printf("Defender: %v\n", res.GetDefender())
	fmt.Printf("Result Counts:\n")
	totalWins := 0
	for k, v := range res.GetResultCounts() {
		fmt.Printf("  Player %v: %v wins\n", k, v)
		totalWins += int(v)
	}
	fmt.Printf("Total Battles: %v\n", totalWins)
}
