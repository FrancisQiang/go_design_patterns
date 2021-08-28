package composite

import (
	"fmt"
)

type Component interface {
	Execute()
}

type Leaf struct {
}

func (l *Leaf) Execute() {
	fmt.Println("leaf execute")
}

type Composite struct {
	children []Component
}

func (composite *Composite) add(component Component) {
	composite.children = append(composite.children, component)
}

func (composite *Composite) remove(i int) {
	composite.children = append(composite.children[:i], composite.children[i+1:]...)
}

func (composite *Composite) Execute() {
	for _, child := range composite.children {
		child.Execute()
	}
	fmt.Println("all execute")
}
