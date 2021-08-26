package factory

import (
	"testing"
)

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("kfc").GenerateFood()
}
