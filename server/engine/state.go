package engine

import (
	"errors"
	"fmt"
	"strings"
)

type Game struct {
	board       [8][8]string
	currentTurn string 
}

func NewGame() *Game {
	g := &Game{
		currentTurn: "white",
	}

	for i := range g.board {
		for j := range g.board[i] {
			g.board[i][j] = "."
		}
	}

	for i := 0; i < 8; i++ {
		g.board[6][i] = "P" // white pawn
		g.board[1][i] = "p" // black pawn
	}

	backRow := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	for i := 0; i < 8; i++ {
		g.board[7][i] = backRow[i]                  
		g.board[0][i] = strings.ToLower(backRow[i]) 
	}

	return g
}

func (g *Game) ApplyMove(from, to string) error {
	fromX, fromY, err := parsePosition(from)
	if err != nil {
		return err
	}
	toX, toY, err := parsePosition(to)
	if err != nil {
		return err
	}

	piece := g.board[fromX][fromY]
	if piece == "." {
		return errors.New("no piece at source")
	}

	if g.currentTurn == "white" && piece != strings.ToUpper(piece) {
		return errors.New("white's turn")
	}
	if g.currentTurn == "black" && piece != strings.ToLower(piece) {
		return errors.New("black's turn")
	}

	if strings.EqualFold(piece, "n") {
		if !isValidKnightMove(fromX, fromY, toX, toY) {
			return errors.New("invalid knight move")
		}
	}
	// temp move logic
	g.board[toX][toY] = piece
	g.board[fromX][fromY] = "."
	fmt.Printf("Moving %s from %s to %s\n", piece, from, to)

	if g.currentTurn == "white" {
		g.currentTurn = "black"
	} else {
		g.currentTurn = "white"
	}

	return nil
}

func (b *Game) PrintBoard() {
	fmt.Println("  a b c d e f g h")
	for i := 7; i >= 0; i-- {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 8; j++ {
			fmt.Printf("%s ", b.board[i][j])
		}
		fmt.Printf("%d\n", i+1)
	}
	fmt.Println("  a b c d e f g h")
}
