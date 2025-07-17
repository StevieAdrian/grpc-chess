package main

import (
	"context"
	"log"
	"time"

	chesspb "grpc-chess/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := chesspb.NewChessServiceClient(conn)

	// temp move from e2 to e4
	req := &chesspb.MoveRequest{
		From:   "e2",
		To:     "e4",
		Player: "white",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.MakeMove(ctx, req)
	if err != nil {
		log.Fatalf("could not make move: %v", err)
	}

	log.Printf("Success: %v, Message: %s", res.Success, res.Message)
}
