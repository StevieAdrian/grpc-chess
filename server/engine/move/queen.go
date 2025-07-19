package move

func QueenMove(fromX, fromY, toX, toY int) bool {
	return BishopMove(fromX, fromY, toX, toY) || RookMove(fromX, fromY, toX, toY)
}
