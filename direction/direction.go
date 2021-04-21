package direction

func RotateLeft(old_direction string) string {
	var new_direction string
	switch old_direction {
	case "NORTH":
		new_direction = "WEST"
	case "EAST":
		new_direction = "NORTH"
	case "SOUTH":
		new_direction = "EAST"
	case "WEST":
		new_direction = "SOUTH"
	}
	return new_direction
}

func RotateRight(old_direction string) string {
	var new_direction string
	switch old_direction {
	case "NORTH":
		new_direction = "EAST"
	case "EAST":
		new_direction = "SOUTH"
	case "SOUTH":
		new_direction = "WEST"
	case "WEST":
		new_direction = "NORTH"
	}
	return new_direction
}

func IsValidDirection(dir string) bool {
	switch dir {
	case
		"NORTH",
		"EAST",
		"SOUTH",
		"WEST":
		return true
	}
	return false
}
