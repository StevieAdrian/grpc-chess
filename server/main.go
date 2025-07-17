package main

import (
	"log"
	"net"

	chesspb "grpc-chess/proto"
	"grpc-chess/server/engine"
	"grpc-chess/server/handler"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	game := engine.NewGame()
	chessServer := &handler.ChessServer{Game: game}

	chesspb.RegisterChessServiceServer(grpcServer, chessServer)

	log.Println("Server started on port 50051")
	game.PrintBoard()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
