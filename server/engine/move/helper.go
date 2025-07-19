package move

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsPathClear(board [8][8]string, fromX, fromY, toX, toY int) bool {
	dx := toX - fromX
	dy := toY - fromY
	stepX, stepY := 0, 0

	if dx != 0 {
		stepX = dx / abs(dx)
	}
	if dy != 0 {
		stepY = dy / abs(dy)
	}

	x, y := fromX+stepX, fromY+stepY
	for x != toX || y != toY {
		if board[x][y] != "." {
			return false
		}
		x += stepX
		y += stepY
	}

	return true
}
