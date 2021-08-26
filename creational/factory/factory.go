package factory

import (
	"fmt"
)

type Restaurant interface {
	GenerateFood()
}

type KFC struct {

}

func (kfc *KFC) GenerateFood() {
	fmt.Println("food has been generated from KFC")
}

type RouXieBao struct {

}

func (rouXieBao *RouXieBao) GenerateFood() {
	fmt.Println("food has been generated from RouXieBao")
}

func NewRestaurant(name string) (restaurant Restaurant) {
	switch name {
	case "kfc":
		return &KFC{}
	case "rxb":
		return &RouXieBao{}
	}
	return nil
}