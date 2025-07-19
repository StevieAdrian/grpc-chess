package move

func BishopMove(fromX, fromY, toX, toY int) bool {
    dx := abs(fromX - toX)
    dy := abs(fromY - toY)
	
    return dx == dy
}