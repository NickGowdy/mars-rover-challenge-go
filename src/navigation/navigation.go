package navigation

import (
	"mars_rover/src/models"
)

func Move(marsRovers []*models.MarsRover, xBoundary int, yBoundary int) {

	for _, marsRover := range marsRovers {
		for _, c := range marsRover.GetInstructions() {
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

	return !(marsRoverCopy.GetX() > xboundary || marsRoverCopy.GetY() > yBoundary)
}

func noPossibleCollision(marsRover *models.MarsRover, marsRovers []*models.MarsRover, instruction string) bool {
	marsRoverCopy := marsRover

	marsRoverCopy.Forward(instruction)

	for _, mr := range marsRovers {
		if mr.GetX() == marsRoverCopy.GetX() && mr.GetY() == marsRover.GetY() {
			return false
		}
	}

	return true
}
