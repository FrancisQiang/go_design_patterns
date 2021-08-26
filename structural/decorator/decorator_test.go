package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	logger := &MyLogger{}
	advanceLogger := &MyAdvanceLogger{logger: logger}
	advanceLogger.LogErrWithColor("hhhh")
}
