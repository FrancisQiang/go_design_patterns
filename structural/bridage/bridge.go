package bridage

import (
	"fmt"
	"sync"
)

type Quantity string

var (
	Little Quantity = "little"
	Middle Quantity = "middle"
	More   Quantity = "more"
)

type MilkTeaShop interface {
	MakeFinishedMilkTea(tea MilkTea, sugar Sugar, ingredients []Ingredients) FinishedMilkTea
}

type ALittleMilkTeaShop struct {
}

func (A ALittleMilkTeaShop) MakeFinishedMilkTea(tea MilkTea, sugar Sugar, ingredients []Ingredients) FinishedMilkTea {
	fmt.Println("welcome to A little tea")
	builder := Builder()
	for _, item := range ingredients {
		builder.AddIngredients(item)
	}
	builder.AddSugar(sugar)
	builder.AddMilkTea(tea)
	return builder.finishedMilkTea
}

type HeyTeaShop struct {
}

func (h HeyTeaShop) MakeFinishedMilkTea(tea MilkTea, sugar Sugar, ingredients []Ingredients) FinishedMilkTea {
	fmt.Println("welcome to Hey tea")
	builder := Builder()
	for _, item := range ingredients {
		builder.AddIngredients(item)
	}
	builder.AddSugar(sugar)
	builder.AddMilkTea(tea)
	return builder.finishedMilkTea
}

type MilkTea interface {
	MakeMilkTea() MilkTea
}

type GreenMilkTea struct {
}

func (g GreenMilkTea) MakeMilkTea() MilkTea {
	return g
}

type BlackMilkTea struct {
}

func (b BlackMilkTea) MakeMilkTea() MilkTea {
	return b
}

type Ingredients interface {
	AddIngredients(quantity Quantity)
}

type Coconut struct {
}

func (c Coconut) AddIngredients(quantity Quantity) {
	fmt.Printf("made %v coconut\n", quantity)
}

type Bubble struct {
}

func (b Bubble) AddIngredients(quantity Quantity) {
	fmt.Printf("made %v bubble\n", quantity)
}

type Sugar interface {
	Percentage() float32
	MakeSugar(percentage float32) Sugar
}

type Toffee struct {
	percentage float32
}

func (t Toffee) Percentage() float32 {
	return t.percentage
}

func (t Toffee) MakeSugar(percentage float32) Sugar {
	t.percentage = percentage
	return t
}

type Caramel struct {
	percentage float32
}

func (c Caramel) Percentage() float32 {
	return c.percentage
}

func (c Caramel) MakeSugar(percentage float32) Sugar {
	c.percentage = percentage
	return c
}

// 如果这样 构建的时候就变成builder模式了
// 但是如果FinishedMilkTea去实现MilkTea，就是桥接模式了
// 但是我反倒觉得这样实现更灵活
type FinishedMilkTea struct {
	ingredients []Ingredients
	sugar       Sugar
	milkTea     MilkTea
}

type FinishedMilkTeaBuilder struct {
	finishedMilkTea FinishedMilkTea
	ingredientsInit *sync.Once
}

func Builder() *FinishedMilkTeaBuilder {
	return &FinishedMilkTeaBuilder{
		ingredientsInit: &sync.Once{},
		finishedMilkTea: FinishedMilkTea{
			ingredients: []Ingredients{},
			sugar:       nil,
			milkTea:     nil,
		},
	}
}

func (builder *FinishedMilkTeaBuilder) AddMilkTea(milkTea MilkTea) {
	builder.finishedMilkTea.milkTea = milkTea
	fmt.Printf("add milkTea %T\n", milkTea)
}

func (builder *FinishedMilkTeaBuilder) AddIngredients(ingredients Ingredients) {
	builder.finishedMilkTea.ingredients = append(builder.finishedMilkTea.ingredients, ingredients)
	fmt.Printf("add ingredients %T\n", ingredients)
}

func (builder *FinishedMilkTeaBuilder) AddSugar(sugar Sugar) {
	builder.finishedMilkTea.sugar = sugar
	fmt.Printf("add %v, %T\n", sugar.Percentage(), sugar)
}

func (builder *FinishedMilkTeaBuilder) Build() FinishedMilkTea {
	return builder.finishedMilkTea
}
