package state

import (
	"testing"
)

func TestState(t *testing.T) {
	ctx := NewContext()
	start := &StartAuthState{}
	start.doAction(ctx)
	twoElement := TwoElementAuthState{}
	twoElement.doAction(ctx)
	success := SuccessAuthState{}
	success.doAction(ctx)
}
