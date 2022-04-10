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
	err := HandleBadDirection(direction)

	if err != nil {
		return nil, err
	}

	m := MarsRover{X: x, Y: y, Direction: direction}
	return &m, nil
}

func Move(marsRover *MarsRover, instructions string) error {
	parsedInstuctions := []byte(instructions)

	for i := 0; i < len(parsedInstuctions); i++ {
		instruction := string(parsedInstuctions[i])

		if instruction == "L" || instruction == "R" {
			Turn(marsRover, instruction)
		} else if instruction == "M" {
			Forward(marsRover, instruction)
		} else {
			return errors.New("Bad input")
		}
	}

	return nil
}

func Turn(marsRover *MarsRover, instruction string) {
	switch marsRover.Direction {
	case 'N':
		if instruction == "L" {
			marsRover.Direction = 'W'
		} else {
			marsRover.Direction = 'E'
		}
	case 'S':
		if instruction == "L" {
			marsRover.Direction = 'E'
		} else {
			marsRover.Direction = 'W'
		}
	case 'E':
		if instruction == "L" {
			marsRover.Direction = 'N'
		} else {
			marsRover.Direction = 'S'
		}
	case 'W':
		if instruction == "L" {
			marsRover.Direction = 'S'
		} else {
			marsRover.Direction = 'N'
		}
	}

}

func Forward(marsRover *MarsRover, instruction string) {
	switch marsRover.Direction {
	case 'N':
		marsRover.Y++
	case 'S':
		marsRover.Y--
	case 'E':
		marsRover.X++
	case 'W':
		marsRover.X--
	}
}

func HandleBadDirection(direction byte) error {
	if direction == 'N' || direction == 'S' || direction == 'E' || direction == 'W' {
		return nil
	}

	return errors.New("Direction invalid")
}
