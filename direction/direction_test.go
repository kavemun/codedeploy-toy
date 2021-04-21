package direction

import (
	"errors"
	"testing"
)

func TestIsValidDirection_WithValidDirection(t *testing.T) {
	expected := true
	actual := IsValidDirection("NORTH")
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestIsNotValidDirection_WithInvalidDirection(t *testing.T) {
	expected := false
	actual := IsValidDirection("HELLO")
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
		err  error
	}{
		{
			name: "rotate left from NORTH",
			in:   "NORTH",
			out:  "WEST",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate left from WEST",
			in:   "WEST",
			out:  "SOUTH",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate left from SOUTH",
			in:   "SOUTH",
			out:  "EAST",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate left from EAST",
			in:   "EAST",
			out:  "NORTH",
			err:  errors.New("wrong direction"),
		},
	}
	for _, test := range tests {
		t.Logf("case: %s", test.name)
		got := RotateLeft(test.in)
		if got != test.out {
			t.Errorf("got %v, expected %v for output", got, test.out)
		}
	}
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
		err  error
	}{
		{
			name: "rotate right from NORTH",
			in:   "NORTH",
			out:  "EAST",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate right from EAST",
			in:   "EAST",
			out:  "SOUTH",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate right from SOUTH",
			in:   "SOUTH",
			out:  "WEST",
			err:  errors.New("wrong direction"),
		},
		{
			name: "rotate right from WEST",
			in:   "WEST",
			out:  "NORTH",
			err:  errors.New("wrong direction"),
		},
	}
	for _, test := range tests {
		t.Logf("case: %s", test.name)
		got := RotateRight(test.in)
		if got != test.out {
			t.Errorf("got %v, expected %v for output", got, test.out)
		}
	}
}
