package main

import (
	"context"
	"fmt"
	"log"
	"net"

	chesspb "grpc-chess/proto"

	"google.golang.org/grpc"
)

type ChessServer struct {
	chesspb.UnimplementedChessServiceServer
}

func (s *ChessServer) MakeMove(ctx context.Context, req *chesspb.MoveRequest) (*chesspb.MoveResponse, error) {
	fmt.Printf("Move from %s to %s by player %s\n", req.From, req.To, req.Player)

	// temp test
	return &chesspb.MoveResponse{
		Success: true,
		Message: "Move accepted",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chesspb.RegisterChessServiceServer(grpcServer, &ChessServer{})

	log.Println("Server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
