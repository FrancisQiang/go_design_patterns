package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	m := &Message{}
	newMessage := m.Clone().(*Message)
	fmt.Println(newMessage)
}
