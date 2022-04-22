package models

import (
	"errors"
	"mars_rover/src/enums"
	"mars_rover/src/validation"
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

	if !validation.HandleBadDirection(direction) {
		return nil, errors.New("bad input for starting direction")
	}
	if !validation.HandleBadInstructions(instructions) {
		return nil, errors.New("bad input for instructions")
	}

	m := MarsRover{x: x, y: y, direction: direction, instructions: instructions}
	return &m, nil
}

func (m *MarsRover) Turn(instruction string) {
	direction := m.Direction()
	switch direction {
	case enums.North:
		if instruction == enums.Left {
			m.direction = enums.West
		} else {
			m.direction = enums.East
		}
	case enums.South:
		if instruction == enums.Left {
			m.direction = enums.East
		} else {
			m.direction = enums.West
		}
	case enums.East:
		if instruction == enums.Left {
			m.direction = enums.North
		} else {
			m.direction = enums.South
		}
	case enums.West:
		if instruction == enums.Left {
			m.direction = enums.South
		} else {
			m.direction = enums.North
		}
	}

}

func (m *MarsRover) Forward(instruction string) {
	switch m.direction {
	case enums.North:
		m.y++
	case enums.South:
		m.y--
	case enums.East:
		m.x++
	case enums.West:
		m.x--
	}
}
