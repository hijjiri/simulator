package main

import (
	"context"
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

	// gRPCリクエスト
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SimulateBattle(ctx, &pb.SimulateRequest{Counts: counts})
	if err != nil {
		log.Fatalf("could not simulate battle: %v", err)
	}

	// レスポンス表示
	log.Printf("Battle ID: %v, Attacker: %v, Defender: %v, Result Counts: %v", res.GetBattleId(), res.GetAttacker(), res.GetDefender(), res.GetResultCounts())
}
