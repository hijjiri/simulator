package service

import (
	"context"
	pb "github.com/hijjiri/simulator/core/go/grpc-server/example"
	"log"
)

type Server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
