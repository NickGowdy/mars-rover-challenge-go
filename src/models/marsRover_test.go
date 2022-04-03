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
