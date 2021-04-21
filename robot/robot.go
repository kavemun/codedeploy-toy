package robot

import (
	"fmt"

	"github.com/kavemun/toy-robot/direction"
	"github.com/kavemun/toy-robot/position"
)

type Robot struct {
	X, Y      int
	Direction string
}

func Place(x int, y int, dir string) (*Robot, error) {
	if !position.IsValidPosition(x, y) {
		return nil, fmt.Errorf("error: position out of bounds")
	}
	if !direction.IsValidDirection(dir) {
		return nil, fmt.Errorf("error: invalid input for direction")
	}
	return &Robot{x, y, dir}, nil
}

func Move(robot *Robot) (*Robot, error) {

	// save original position
	orig_x := robot.X
	orig_y := robot.Y

	switch robot.Direction {
	case "NORTH":
		robot.Y += 1
	case "EAST":
		robot.X += 1
	case "SOUTH":
		robot.Y -= 1
	case "WEST":
		robot.X -= 1
	}
	if !position.IsValidPosition(robot.X, robot.Y) {
		// cancel robot movement
		return &Robot{orig_x, orig_y, robot.Direction}, fmt.Errorf("robot trying to move out of bound")
	}
	return robot, nil
}

func RotateLeft(robot *Robot) *Robot {
	return &Robot{robot.X, robot.Y, direction.RotateLeft(robot.Direction)}
}

func RotateRight(robot *Robot) *Robot {
	return &Robot{robot.X, robot.Y, direction.RotateRight(robot.Direction)}
}

func ReportPosition(robot *Robot) string {
	return fmt.Sprintf("Output: %d,%d,%s", robot.X, robot.Y, robot.Direction)
}
