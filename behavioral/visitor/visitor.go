package visitor

import (
	"fmt"
)

type Visitor interface {
	VisitCity()
	VisitVillage()
}

type NormalVisitor struct {
}

func (n NormalVisitor) VisitCity() {
	fmt.Println("city is bad")
}

func (n NormalVisitor) VisitVillage() {
	fmt.Println("village is bad")
}

type InspectVisitor struct {
}

func (i InspectVisitor) VisitCity() {
	fmt.Println("city is good")
}

func (i InspectVisitor) VisitVillage() {
	fmt.Println("village is good")
}

type Country interface {
	Accept(visitor Visitor)
}

type CountryA struct {
	
}

func (c CountryA) Accept(visitor Visitor) {
	visitor.VisitCity()
}

type CountryB struct {

}

func (c CountryB) Accept(visitor Visitor) {
	visitor.VisitVillage()
}
