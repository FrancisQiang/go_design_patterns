package factory

import (
	"fmt"
)

type StableFood interface {
	Eat()
}

type SideDish interface {
	Taste()
}

type Rise struct {
}

type Tomato struct {
}

type Noodle struct {
}

type Egg struct {
}

func (rise *Rise) Eat() {
	fmt.Println("eat rise")
}

func (tomato *Tomato) Taste() {
	fmt.Println("taste tomato")
}

func (noodle *Noodle) Eat() {
	fmt.Println("eat Noodle")
}

func (egg *Egg) Taste() {
	fmt.Println("taste egg")
}

type FoodFactory interface {
	CreateStableFood() StableFood
	CreateSideDish() SideDish
}

type TomatoRiseFoodFactory struct {
}

func (t *TomatoRiseFoodFactory) CreateStableFood() StableFood {
	return &Rise{}
}

func (t *TomatoRiseFoodFactory) CreateSideDish() SideDish {
	return &Tomato{}
}

type EggNoodleFoodFactory struct {
}

func (e *EggNoodleFoodFactory) CreateStableFood() StableFood {
	return &Noodle{}
}

func (e *EggNoodleFoodFactory) CreateSideDish() SideDish {
	return &Egg{}
}

func NewFoodFactory(factoryName string) FoodFactory {
	switch factoryName {
	case "tomatoRise":
		return &TomatoRiseFoodFactory{}
	case "eggNoodle":
		return &EggNoodleFoodFactory{}

	}
	return nil
}
