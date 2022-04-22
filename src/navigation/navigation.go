package navigation

import (
	"mars_rover/src/models"
)

func Move(marsRovers []*models.MarsRover, xBoundary int, yBoundary int) {

	for _, marsRover := range marsRovers {
		for _, c := range marsRover.Instructions() {
			parsedInstruction := string(c)

			if parsedInstruction == "L" || parsedInstruction == "R" {
				marsRover.Turn(string(c))
			} else if parsedInstruction == "M" &&
				noPossibleCollision(marsRover, marsRovers, parsedInstruction) &&
				isInsideBoundary(marsRover, parsedInstruction, xBoundary, yBoundary) {
				marsRover.Forward(parsedInstruction)
			}
		}
	}
}

func isInsideBoundary(marsRover *models.MarsRover, instruction string, xboundary int, yBoundary int) bool {
	marsRoverCopy := marsRover
	marsRoverCopy.Forward(instruction)

	return !(marsRoverCopy.X() > xboundary || marsRoverCopy.Y() > yBoundary)
}

func noPossibleCollision(marsRover *models.MarsRover, marsRovers []*models.MarsRover, instruction string) bool {
	marsRoverCopy := marsRover

	marsRoverCopy.Forward(instruction)

	for _, mr := range marsRovers {
		if mr.X() == marsRoverCopy.X() && mr.Y() == marsRover.Y() {
			return false
		}
	}

	return true
}
