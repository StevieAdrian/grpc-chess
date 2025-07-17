package main

import (
	"log"

	"grpc-chess/client/engine"
	chesspb "grpc-chess/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect: %v", err)
	}
	defer conn.Close()

	client := chesspb.NewChessServiceClient(conn)

	for {
		player, from, to, ok := engine.ReadMoveInput()
		if !ok {
			break
		}
		engine.MakeMove(client, player, from, to)
	}
}
