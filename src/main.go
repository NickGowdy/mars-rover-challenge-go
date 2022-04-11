package main

import (
	"mars_rover/src/models"
	"mars_rover/src/navigation"
)

func main() {
	rover1, _ := models.NewRover(1, 2, "N", "LMLMLMLMM")
	rover2, _ := models.NewRover(3, 3, "E", "MMRMMRMRRM")

	navigation.Move([]*models.MarsRover{rover1, rover2}, 5, 5)
}
