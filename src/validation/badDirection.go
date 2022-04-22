package validation

import "mars_rover/src/enums"

func HandleBadDirection(direction string) bool {
	return direction == enums.North || direction == enums.South || direction == enums.East || direction == enums.West
}

func HandleBadInstructions(instructions string) bool {
	isValid := true

	for _, c := range instructions {
		if string(c) == enums.Left || string(c) == enums.Right || string(c) == enums.Move {
			continue
		} else {
			isValid = false
			break
		}

	}
	return isValid
}
