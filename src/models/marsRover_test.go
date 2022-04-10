package models

import "testing"

func TestNewMarsRover(t *testing.T) {
	rover, _ := NewRover(1, 2, 'N')

	if rover == nil {
		t.Error("Mars Rover can't be nil")
	}
}

func TestNewMarsRoverDirection(t *testing.T) {
	_, err := NewRover(1, 2, 'Z')

	if err == nil {
		t.Error("Mars Rover directions can only be N, S, E, W")
	}
}

func TestMarsRoverMove(t *testing.T) {
	rover, _ := NewRover(1, 2, 'N')

	instructions := "LMLMLMLMM"

	Move(rover, instructions)

	if rover.X != 1 {
		t.Error("Rover X position should be 1")
	}

	if rover.Y != 3 {
		t.Error("Rover Y position should be 1")
	}
}
