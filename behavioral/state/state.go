package state

import (
	"fmt"
)

type Context struct {
	state State
}

func NewContext() Context {
	return Context{}
}

func (ctx Context) SetState(state State) {
	ctx.state = state
}

func (ctx Context) GetState() State {
	return ctx.state
}

type State interface {
	doAction(ctx Context)
}

type StartAuthState struct {

}

func (s StartAuthState) doAction(ctx Context) {
	fmt.Println("auth start")
	ctx.SetState(s)
}

type TwoElementAuthState struct {

}

func (t TwoElementAuthState) doAction(ctx Context) {
	fmt.Println("two element has been verified")
	ctx.SetState(t)
}

type SuccessAuthState struct {

}

func (s SuccessAuthState) doAction(ctx Context) {
	fmt.Println("face verified success")
	ctx.SetState(s)
}

type FailedAuthState struct {

}

func (f FailedAuthState) doAction(ctx Context) {
	fmt.Println("face verified failed")
	ctx.SetState(f)
}




