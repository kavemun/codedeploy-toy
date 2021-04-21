package robot

import (
	"errors"
	"testing"
)

type input struct {
	x         int
	y         int
	direction string
}

func TestPlace(t *testing.T) {

	tests := []struct {
		name string
		in   input
		out  *Robot
		err  error
	}{
		{
			name: "Place OK",
			in:   input{0, 0, "NORTH"},
			out:  &Robot{0, 0, "NORTH"},
			err:  nil,
		},
	}

	for _, test := range tests {
		t.Logf("case: %s", test.name)
		actual, _ := Place(test.in.x, test.in.y, test.in.direction)
		expected := test.out
		if actual.X != expected.X || actual.Y != expected.Y || actual.Direction != expected.Direction {
			t.Errorf("got %v, want %v for output", actual, expected)
		}
	}

}

func TestPlace_WithInvalid(t *testing.T) {

	tests := []struct {
		name string
		in   input
		out  *Robot
		err  error
	}{
		{
			name: "Place with invalid position",
			in:   input{0, 6, "NORTH"},
			out:  nil,
			err:  nil,
		},
		{
			name: "Place with invalid direction",
			in:   input{0, 0, "NOWHERE"},
			out:  nil,
			err:  nil,
		},
	}

	for _, test := range tests {
		t.Logf("case: %s", test.name)
		actual, _ := Place(test.in.x, test.in.y, test.in.direction)
		expected := test.out
		if actual != expected {
			t.Errorf("got %v, want %v for output", actual, expected)
		}
	}

}

func TestMove(t *testing.T) {

	tests := []struct {
		name string
		in   *Robot
		out  *Robot
		err  error
	}{
		{
			name: "Move North",
			in:   &Robot{0, 0, "NORTH"},
			out:  &Robot{0, 1, "NORTH"},
			err:  nil,
		},
		{
			name: "Move South",
			in:   &Robot{0, 1, "SOUTH"},
			out:  &Robot{0, 0, "SOUTH"},
			err:  nil,
		},
		{
			name: "Move EAST",
			in:   &Robot{0, 1, "EAST"},
			out:  &Robot{1, 1, "EAST"},
			err:  nil,
		},
		{
			name: "Move WEST",
			in:   &Robot{4, 1, "WEST"},
			out:  &Robot{3, 1, "WEST"},
			err:  nil,
		},
		{
			name: "Move Past Edge",
			in:   &Robot{0, 3, "WEST"},
			out:  &Robot{0, 3, "WEST"},
			err:  errors.New("Error: Robot trying to move out of bound."),
		},
	}

	for _, test := range tests {
		t.Logf("case: %s", test.name)
		actual, _ := Move(test.in)
		expected := test.out
		// if err != nil {
		// 	t.Errorf("Error %v", err)
		// }
		if actual.X != expected.X || actual.Y != expected.Y || actual.Direction != expected.Direction {
			t.Errorf("got %v, want %v for output", actual, expected)
		}
	}

}

func TestReport(t *testing.T) {
	expected := "Output: 4,4,SOUTH"
	robot, err := Place(4, 4, "SOUTH")
	if err != nil {
		t.Errorf("Robot failed to place")
	}
	actual := ReportPosition(robot)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestLeft(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = RotateLeft(robot)
	if robot.Direction != "WEST" {
		t.Errorf("Robot has not turned left")
	}
}

func TestRight(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = RotateRight(robot)
	if robot.Direction != "EAST" {
		t.Errorf("Robot has not turned right")
	}
}
