package models

type MarsRover struct {
	x         int
	y         int
	direction string
}

func NewRover(x int, y int, direction string) *MarsRover {
	m := MarsRover{x: x, y: y, direction: direction}
	return &m
}
