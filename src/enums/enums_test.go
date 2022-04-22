package enums_test

import (
	"mars_rover/src/enums"
	"testing"
)

func TestEnumsReturnCorrectValue(t *testing.T) {
	if enums.North != "N" {
		t.Error(enums.North + " should equal E")
	}

	if enums.South != "S" {
		t.Error(enums.South + " should equal S")
	}

	if enums.West != "W" {
		t.Error(enums.West + " should equal W")
	}

	if enums.East != "E" {
		t.Error(enums.East + " should equal E")
	}

	if enums.Right != "R" {
		t.Error(enums.Right + " should equal R")
	}

	if enums.Left != "L" {
		t.Error(enums.Left + " should equal L")
	}
}
