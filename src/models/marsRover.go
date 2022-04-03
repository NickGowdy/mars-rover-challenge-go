package models

import (
	"errors"
)

type MarsRover struct {
	X         int
	Y         int
	Direction byte
}

func NewRover(x int, y int, direction byte) (*MarsRover, error) {
	err := HandleBadInput(direction)

	if err != nil {
		return nil, err
	}

	m := MarsRover{X: x, Y: y, Direction: direction}
	return &m, nil
}

func HandleBadInput(direction byte) error {
	if direction == 'N' || direction == 'S' || direction == 'E' || direction == 'W' {
		return nil
	}

	return errors.New("Direction invalid")
}
