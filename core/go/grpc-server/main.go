package main

import (
	"log"
	"net"
	"os"

	pb_example "github.com/hijjiri/simulator/core/go/grpc-server/example"
	example_service "github.com/hijjiri/simulator/core/go/grpc-server/example/service"
	pb_janken "github.com/hijjiri/simulator/core/go/grpc-server/janken"
	janken_service "github.com/hijjiri/simulator/core/go/grpc-server/janken/service"
	pb_simulator "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	simulator_service "github.com/hijjiri/simulator/core/go/grpc-server/simulator/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb_example.RegisterExampleServiceServer(s, &example_service.Server{})
	pb_janken.RegisterJankenServiceServer(s, &janken_service.Server{})
	pb_simulator.RegisterSimulatorServiceServer(s, &simulator_service.Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
