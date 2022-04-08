package models

type Surface struct {
	X int
	Y int
}

func CreateSurface(x int, y int) (*Surface, error) {
	return &Surface{}, nil
}
