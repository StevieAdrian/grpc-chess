package move

func KingMove(fromX, fromY, toX, toY int) bool {
	dx := abs(fromX - toX)
	dy := abs(fromY - toY)

	return dx <= 1 && dy <= 1
}
