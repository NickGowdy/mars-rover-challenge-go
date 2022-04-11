package models

import (
	"errors"
)

type MarsRover struct {
	X            int
	Y            int
	Direction    string
	Instructions string
}

func NewRover(x int, y int, direction string, instructions string) (*MarsRover, error) {

	if !handleBadDirection(direction) {
		return nil, errors.New("bad input for starting direction")
	}
	if !handleBadInstructions(instructions) {
		return nil, errors.New("bad input for instructions")
	}

	m := MarsRover{X: x, Y: y, Direction: direction, Instructions: instructions}
	return &m, nil
}

func handleBadDirection(direction string) bool {
	return direction == "N" || direction == "S" || direction == "E" || direction == "W"
}

func handleBadInstructions(instructions string) bool {
	isValid := true

	for _, c := range instructions {
		if string(c) == "L" || string(c) == "R" || string(c) == "M" {
			continue
		} else {
			isValid = false
			break
		}

	}
	return isValid
}
