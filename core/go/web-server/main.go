package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/hijjiri/simulator/core/go/grpc-server/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)

	r := gin.Default()

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": res.GetMessage()})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
