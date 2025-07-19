package main

import (
	"fmt"
	"strings"
)

func squareToIndex(sq string) (row, col int, err error) {
	if len(sq) != 2 {
		return 0, 0, fmt.Errorf("bad square: %q", sq)
	}
	file := strings.ToLower(string(sq[0])) // a-h
	rank := string(sq[1])                  // 1-8

	if file[0] < 'a' || file[0] > 'h' {
		return 0, 0, fmt.Errorf("invalid file: %s", file)
	}
	if rank[0] < '1' || rank[0] > '8' {
		return 0, 0, fmt.Errorf("invalid rank: %s", rank)
	}

	col = int(file[0] - 'a')        // a=0 ... h=7
	boardRank := int(rank[0] - '0') // '1' -> 1
	row = 8 - boardRank             // rank 8 -> row0, rank1 -> row7
	return row, col, nil
}
