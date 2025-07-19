package move

func KnightMove(fromX, fromY, toX, toY int) bool {
	dx := abs(fromX - toX)
	dy := abs(fromY - toY)

	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
}
