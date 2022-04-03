package navigation

import (
	"testing"
)

func TestNavigation(t *testing.T) {
	rover := newRover(1, 2, "N")

	if rover == nil {
		t.Error("Mars Rover can't be nil")
	}
}
