package handler

import (
	"context"
	"fmt"
	chesspb "grpc-chess/proto"
	"grpc-chess/server/engine"
)

type ChessServer struct {
	chesspb.UnimplementedChessServiceServer
	Game *engine.Game
}

func (s *ChessServer) MakeMove(ctx context.Context, req *chesspb.MoveRequest) (*chesspb.MoveResponse, error) {
	from := req.From
	to := req.To
	player := req.Player

	fmt.Printf("Req: %s -> %s (%s)\n", from, to, player)

	if err := s.Game.ApplyMove(from, to); err != nil {
		return &chesspb.MoveResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	s.Game.PrintBoard()

	return &chesspb.MoveResponse{
		Success: true,
		Message: "Move applied",
	}, nil
}
