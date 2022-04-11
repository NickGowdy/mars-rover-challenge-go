package models

import (
	"errors"
)

type MarsRover struct {
	x, y                    int
	direction, instructions string
}

func (marsRover *MarsRover) GetX() int {
	return marsRover.x
}

func (marsRover *MarsRover) GetY() int {
	return marsRover.y
}

func (marsRover *MarsRover) GetInstructions() string {
	return marsRover.instructions
}

func (marsRover *MarsRover) GetDirection() string {
	return marsRover.direction
}

func NewRover(x int, y int, direction string, instructions string) (*MarsRover, error) {

	if !handleBadDirection(direction) {
		return nil, errors.New("bad input for starting direction")
	}
	if !handleBadInstructions(instructions) {
		return nil, errors.New("bad input for instructions")
	}

	m := MarsRover{x: x, y: y, direction: direction, instructions: instructions}
	return &m, nil
}

func (marsRover *MarsRover) Turn(instruction string) {
	direction := marsRover.GetDirection()
	switch direction {
	case "N":
		if instruction == "L" {
			marsRover.direction = "W"
		} else {
			marsRover.direction = "E"
		}
	case "S":
		if instruction == "L" {
			marsRover.direction = "E"
		} else {
			marsRover.direction = "W"
		}
	case "E":
		if instruction == "L" {
			marsRover.direction = "N"
		} else {
			marsRover.direction = "S"
		}
	case "W":
		if instruction == "L" {
			marsRover.direction = "S"
		} else {
			marsRover.direction = "N"
		}
	}

}

func (marsRover *MarsRover) Forward(instruction string) {
	switch marsRover.direction {
	case "N":
		marsRover.y++
	case "S":
		marsRover.y--
	case "E":
		marsRover.x++
	case "W":
		marsRover.x--
	}
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
