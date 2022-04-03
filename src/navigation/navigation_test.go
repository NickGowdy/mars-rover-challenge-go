package navigation

import (
	"mars_rover/src/models"
	"testing"
)

func TestNavigation(t *testing.T) {
	rover := models.NewRover(1, 2, "N")

	if rover == nil {
		t.Error("Mars Rover can't be nil")
	}
}
