package models

import (
	"errors"
)

type MarsRover struct {
	x, y         int
	direction    string
	instructions string
}

func (m *MarsRover) X() int {
	return m.x
}

func (m *MarsRover) Y() int {
	return m.y
}

func (m *MarsRover) Instructions() string {
	return m.instructions
}

func (m *MarsRover) Direction() string {
	return m.direction
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

func (m *MarsRover) Turn(instruction string) {
	direction := m.Direction()
	switch direction {
	case "N":
		if instruction == "L" {
			m.direction = "W"
		} else {
			m.direction = "E"
		}
	case "S":
		if instruction == "L" {
			m.direction = "E"
		} else {
			m.direction = "W"
		}
	case "E":
		if instruction == "L" {
			m.direction = "N"
		} else {
			m.direction = "S"
		}
	case "W":
		if instruction == "L" {
			m.direction = "S"
		} else {
			m.direction = "N"
		}
	}

}

func (m *MarsRover) Forward(instruction string) {
	switch m.direction {
	case "N":
		m.y++
	case "S":
		m.y--
	case "E":
		m.x++
	case "W":
		m.x--
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
