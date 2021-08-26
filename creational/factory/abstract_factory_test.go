package factory

import (
	"testing"
)

func TestNewFoodFactory(*testing.T) {
	foodFactory := NewFoodFactory("eggNoodle")
	foodFactory.CreateStableFood().Eat()
	foodFactory.CreateSideDish().Taste()
}
