package engine

import "errors"

func parsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, errors.New("invalid position format")
	}

	col := int(pos[0] - 'a') 
	row := int(pos[1] - '1') 

	if col < 0 || col > 7 || row < 0 || row > 7 {
		return 0, 0, errors.New("position out of bounds")
	}

	return row, col, nil
}
