package move

func RookMove(fromX, fromY, toX, toY int) bool {
    return fromX == toX || fromY == toY
}
