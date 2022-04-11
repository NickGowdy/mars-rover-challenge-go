package navigation

import (
	"mars_rover/src/models"
	"testing"
)

func TestNavigation(t *testing.T) {

	rover, _ := models.NewRover(1, 2, "N", "LMLMLMLMM")

	Move([]*models.MarsRover{rover})

	if rover.X != 1 {
		t.Error("Rover X coordinate should be 1")
	}

	if rover.Y != 3 {
		t.Error("Rover Y coordinate should be 1")
	}

	if rover.Direction != "N" {
		t.Error("Rover should be facing direction N")
	}
}
