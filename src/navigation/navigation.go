package navigation

import (
	"mars_rover/src/models"
)

func Move(marsRovers []*models.MarsRover) {

	for _, marsRover := range marsRovers {
		for _, c := range marsRover.Instructions {
			parsedInstruction := string(c)

			if parsedInstruction == "L" || parsedInstruction == "R" {
				turn(marsRover, string(c))
			} else if parsedInstruction == "M" && noPossibleCollision(marsRover, marsRovers, parsedInstruction) {
				forward(marsRover, string(c))
			}
		}
	}
}

func noPossibleCollision(marsRover *models.MarsRover, marsRovers []*models.MarsRover, instruction string) bool {
	marsRoverCopy := marsRover

	forward(marsRoverCopy, instruction)

	for _, mr := range marsRovers {
		if mr.X == marsRoverCopy.X && mr.Y == marsRover.Y {
			return false
		}
	}

	return true
}

func turn(marsRover *models.MarsRover, instruction string) {
	switch marsRover.Direction {
	case "N":
		if instruction == "L" {
			marsRover.Direction = "W"
		} else {
			marsRover.Direction = "E"
		}
	case "S":
		if instruction == "L" {
			marsRover.Direction = "E"
		} else {
			marsRover.Direction = "W"
		}
	case "E":
		if instruction == "L" {
			marsRover.Direction = "N"
		} else {
			marsRover.Direction = "S"
		}
	case "W":
		if instruction == "L" {
			marsRover.Direction = "S"
		} else {
			marsRover.Direction = "N"
		}
	}

}

func forward(marsRover *models.MarsRover, instruction string) {
	switch marsRover.Direction {
	case "N":
		marsRover.Y++
	case "S":
		marsRover.Y--
	case "E":
		marsRover.X++
	case "W":
		marsRover.X--
	}
}
