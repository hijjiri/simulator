package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hijjiri/simulator/core/go/grpc-server/aura"
	aura_repository "github.com/hijjiri/simulator/core/go/grpc-server/aura/repository"
	"github.com/hijjiri/simulator/core/go/grpc-server/battle"
	battle_repository "github.com/hijjiri/simulator/core/go/grpc-server/battle/repository"
	pb_example "github.com/hijjiri/simulator/core/go/grpc-server/example"
	example_service "github.com/hijjiri/simulator/core/go/grpc-server/example/service"
	pb_janken "github.com/hijjiri/simulator/core/go/grpc-server/janken"
	janken_service "github.com/hijjiri/simulator/core/go/grpc-server/janken/service"
	pb_simulator "github.com/hijjiri/simulator/core/go/grpc-server/simulator"
	simulator_service "github.com/hijjiri/simulator/core/go/grpc-server/simulator/service"
	"github.com/hijjiri/simulator/core/go/grpc-server/user"
	user_repository "github.com/hijjiri/simulator/core/go/grpc-server/user/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initializeSkills() {
	skills := map[uint32]*battle.Skill{
		1: {SkillId: 1},
		2: {SkillId: 2},
		// 必要なスキルをすべて追加
	}
	battleRepo := battle_repository.NewInMemoryRepository()
	battleRepo.InitializeSkills(skills)
	battle.RegisterRepository(battleRepo)
}

func initializeAuraTypes() {
	auraTypes := map[uint32]*aura.AuraTypeMaster{
		1: {AuraType: 1},
		2: {AuraType: 2},
		// 必要なオーラタイプをすべて追加
	}
	auraRepo := aura_repository.NewInMemoryRepository()
	auraRepo.InitializeAuraTypes(auraTypes)
	aura.RegisterRepository(auraRepo)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register in-memory user repository
	userRepo := user_repository.NewInMemoryRepository()
	user.RegisterRepository(userRepo)

	// Initialize repositories
	initializeSkills()
	initializeAuraTypes()

	s := grpc.NewServer()
	pb_example.RegisterExampleServiceServer(s, &example_service.Server{})
	pb_janken.RegisterJankenServiceServer(s, &janken_service.Server{})
	pb_simulator.RegisterSimulatorServiceServer(s, &simulator_service.Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	// Graceful shutdown
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Catch shutdown signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("stopping server...")
	s.GracefulStop()
	log.Println("server stopped")
}
