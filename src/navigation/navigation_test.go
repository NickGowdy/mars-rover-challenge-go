package navigation

import (
	"mars_rover/src/models"
	"testing"
)

func TestNavigation(t *testing.T) {

	rover1, _ := models.NewRover(1, 2, "N", "LMLMLMLMM")
	rover2, _ := models.NewRover(3, 3, "E", "MMRMMRMRRM")

	Move([]*models.MarsRover{rover1, rover2}, 5, 5)

	if rover1.X != 1 {
		t.Error("Rover X coordinate should be 1")
	}

	if rover1.Y != 3 {
		t.Error("Rover Y coordinate should be 1")
	}

	if rover1.Direction != "N" {
		t.Error("Rover should be facing direction N")
	}

	if rover2.X != 5 {
		t.Error("Rover X coordinate should be 1")
	}

	if rover2.Y != 1 {
		t.Error("Rover Y coordinate should be 1")
	}

	if rover2.Direction != "E" {
		t.Error("Rover should be facing direction N")
	}
}
