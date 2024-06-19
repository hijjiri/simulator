package main

import (
    "log"
    "net"
    "os"

    "google.golang.org/grpc"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "50051" // デフォルトポート
    }
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    // gRPCサーバーの登録コードをここに追加
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
