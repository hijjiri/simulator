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

// 定数とグローバル変数の定義
const (
	workerNum = 10  // ワーカーの数
	queueSize = 100 // リクエストキューのサイズ
)

var requestQueue chan func() // リクエストキューを保持するチャネル

// init関数
func init() {
	requestQueue = make(chan func(), queueSize) // リクエストキューの初期化
	for i := 0; i < workerNum; i++ {            // ワーカーの数だけゴルーチンを起動
		go worker() // 各ワーカーゴルーチンを起動
	}
}

// ワーカーファンクション
func worker() {
	for req := range requestQueue { // リクエストキューからリクエストを受け取る
		req() // リクエストを処理
	}
}

// Server構造体の定義
type Server struct {
	pb.UnimplementedSimulatorServiceServer
}

// SimulateBattle関数
func (s *Server) SimulateBattle(ctx context.Context, in *pb.SimulateRequest) (*pb.SimulateResponse, error) {
	resultCh := make(chan *pb.SimulateResponse, 1) // 結果を受け取るチャネル
	errorCh := make(chan error, 1)                 // エラーを受け取るチャネル

	var wg sync.WaitGroup
	wg.Add(1) // ワーカーグループに1つのワーカーを追加

	requestQueue <- func() { // リクエストキューにリクエストを追加
		defer wg.Done() // 関数の終了時にワーカーグループから1つのワーカーを削除

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
		wg.Wait() // ワーカーグループがすべて完了するのを待つ
		close(resultCh)
		close(errorCh)
	}()

	// select文で結果またはエラーのいずれかを待つ
	select {
	case res := <-resultCh:
		return res, nil
	case err := <-errorCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
