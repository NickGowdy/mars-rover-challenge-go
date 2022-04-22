package validation

import (
	"mars_rover/src/enums"
	"testing"
)

func TestHandleBadDirectionReturnsCorrectValue(t *testing.T) {
	if HandleBadDirection(enums.North) != true {
		t.Error(enums.North + " Should return true")
	}

	if HandleBadDirection(enums.South) != true {
		t.Error(enums.South + " Should return true")
	}

	if HandleBadDirection(enums.East) != true {
		t.Error(enums.East + " Should return true")
	}

	if HandleBadDirection(enums.West) != true {
		t.Error(enums.West + " Should return true")
	}

	if HandleBadDirection(enums.Right) == true {
		t.Error(enums.Right + " Should return false")
	}

	if HandleBadDirection(enums.Left) == true {
		t.Error(enums.Left + " Should return false")
	}

	if HandleBadDirection(enums.Move) == true {
		t.Error(enums.Move + " Should return false")
	}
}

func TestHandleBadInstructionsReturnsCorrectValue(t *testing.T) {
	if HandleBadInstructions(enums.Left) != true {
		t.Error(enums.Left + " Should return true")
	}

	if HandleBadInstructions(enums.Right) != true {
		t.Error(enums.Right + " Should return true")
	}

	if HandleBadInstructions(enums.Move) != true {
		t.Error(enums.Move + " Should return true")
	}

	if HandleBadInstructions(enums.North) == true {
		t.Error(enums.North + " Should return false")
	}

	if HandleBadInstructions(enums.South) == true {
		t.Error(enums.South + " Should return false")
	}

	if HandleBadInstructions(enums.East) == true {
		t.Error(enums.East + " Should return false")
	}

	if HandleBadInstructions(enums.West) == true {
		t.Error(enums.West + " Should return false")
	}
}
