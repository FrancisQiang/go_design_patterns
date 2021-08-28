package facade

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Circle struct {

}

func (c Circle) Draw() {
	fmt.Println("Draw Circle")
}

func (r Rectangle) Draw() {
	fmt.Println("Draw Rectangle")
}

func (t Triangle) Draw() {
	fmt.Println("Draw Triangle")
}

type Rectangle struct {

}

type Triangle struct {

}

type ShapeMaker struct {
	circle Circle
	rectangle Rectangle
	triangle Triangle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle: Circle{},
		rectangle: Rectangle{},
		triangle: Triangle{},
	}
}

func (shapeMaker *ShapeMaker) drawCircle() {
	shapeMaker.circle.Draw()
}

func (shapeMaker *ShapeMaker) drawRectangle() {
	shapeMaker.rectangle.Draw()
}

func (shapeMaker *ShapeMaker) drawTriangle() {
	shapeMaker.triangle.Draw()
}
