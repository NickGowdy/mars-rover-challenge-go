package navigation

import (
	"fmt"
	"mars_rover/src/models"
	"testing"
)

func TestNavigation(t *testing.T) {

	rover1, _ := models.NewRover(1, 2, "N", "LMLMLMLMM")
	rover2, _ := models.NewRover(3, 3, "E", "MMRMMRMRRM")

	Move([]*models.MarsRover{rover1, rover2}, 5, 5)

	fmt.Println(rover1.GetX())
	fmt.Println(rover1.GetY())

	if rover1.GetX() != 1 {
		t.Error("Rover X coordinate should be 1")
	}

	if rover1.GetY() != 3 {
		t.Error("Rover Y coordinate should be 1")
	}

	if rover1.GetDirection() != "N" {
		t.Error("Rover should be facing direction N")
	}

	if rover2.GetX() != 5 {
		t.Error("Rover X coordinate should be 1")
	}

	if rover2.GetY() != 1 {
		t.Error("Rover Y coordinate should be 1")
	}

	if rover2.GetDirection() != "E" {
		t.Error("Rover should be facing direction N")
	}
}
