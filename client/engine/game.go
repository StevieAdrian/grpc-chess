package engine

import (
	"context"
	"fmt"
	"time"

	chesspb "grpc-chess/proto"
)

func MakeMove(client chesspb.ChessServiceClient, player, from, to string) {
	req := &chesspb.MoveRequest{
		Player: player,
		From:   from,
		To:     to,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.MakeMove(ctx, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Success: %v, Message: %s\n", res.Success, res.Message)
}
