package position

const MAX_X, MAX_Y int = 4, 4

func IsValidPosition(x int, y int) bool {
	if x > MAX_X || x < 0 || y > MAX_Y || y < 0 {
		return false
	}
	return true
}
