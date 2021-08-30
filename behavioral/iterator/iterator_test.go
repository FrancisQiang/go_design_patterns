package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T)  {
	e1 := &Element{
		value: "a",
	}
	e2 := &Element{
		value: "b",
	}
	e3 := &Element{
		value: "c",
	}
	ec := ElementCollection{elements: []*Element{e1, e2, e3}}
	iterator := ec.Iterator()
	for iterator.HasNext() {
		e := iterator.Next()
		fmt.Println(e.value)
	}
}
