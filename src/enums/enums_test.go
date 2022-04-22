package enums

import (
	"testing"
)

func TestEnumsReturnCorrectValue(t *testing.T) {
	if North != "N" {
		t.Error(North + " should equal E")
	}

	if South != "S" {
		t.Error(South + " should equal S")
	}

	if West != "W" {
		t.Error(West + " should equal W")
	}

	if East != "E" {
		t.Error(East + " should equal E")
	}

	if Right != "R" {
		t.Error(Right + " should equal R")
	}

	if Left != "L" {
		t.Error(Left + " should equal L")
	}
}
