package models

import "testing"

func TestNewMarsRover(t *testing.T) {
	rover, _ := NewRover(1, 2, "N", "L")

	if rover == nil {
		t.Error("Mars Rover can't be nil")
	}
}

func TestNewMarsRoverDirection(t *testing.T) {
	_, err := NewRover(1, 2, "Z", "LMLMLMLMM")

	if err == nil {
		t.Error("Mars Rover directions can only be N, S, E, W")
	}
}

func TestNewMarsRoverInstructions(t *testing.T) {
	_, err := NewRover(1, 2, "Z", "123445")

	if err == nil {
		t.Error("Mars Rover directions can only contain L, R, M")
	}
}
