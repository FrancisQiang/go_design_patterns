package template

import (
	"testing"
)

func TestTemplate(*testing.T) {
	worker := Daily{
		worker: Coder{},
	}
	worker.Live()
}