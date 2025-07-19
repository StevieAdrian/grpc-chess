package move

func PawnMove(fromX, fromY, toX, toY int, isWhite bool, targetPiece string) bool {
	direction := -1
	startRow := 6

	if !isWhite {
		direction = 1
		startRow = 1
	}

	dx := toX - fromX
	dy := abs(toY - fromY)

	if dy == 0 && targetPiece == "." {
		if dx == direction {
			return true
		}
		if fromX == startRow && dx == 2*direction {
			return true
		}
	}

	if dy == 1 && dx == direction && targetPiece != "." {
		return true
	}

	return false
}
