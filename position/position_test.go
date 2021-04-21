package position

import "testing"

func TestIsValidPosition_WithCorrectPlacing(t *testing.T) {
	expected := true
	actual := IsValidPosition(0, 1)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestIsValidPosition_WithInvalidPlacing(t *testing.T) {
	expected := false
	actual := IsValidPosition(6, 1)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
