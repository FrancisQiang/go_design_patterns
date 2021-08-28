package bridage

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	aLittleMilkTeaShop := &ALittleMilkTeaShop{}
	sugar := Toffee{}.MakeSugar(0.5)
	tea := aLittleMilkTeaShop.MakeFinishedMilkTea(BlackMilkTea{}, sugar, []Ingredients{Bubble{}, Coconut{}})
	fmt.Printf("made a milkTea %v", tea)
}
