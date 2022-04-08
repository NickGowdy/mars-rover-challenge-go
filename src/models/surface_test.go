package models

import "testing"

func CreateTest(t *testing.T) {
	x := 5
	y := 5
	surface, _ := CreateSurface(x, y)

	if surface == nil {
		t.Error("Surface can't be nil")
	}
}
