package responsebility

import (
	"fmt"
)

type Validator interface {
	Validate(params string)
	SetNext(validator *Validator)
	Next(params string)
}

type TwoElementValidator struct {
	next *Validator
}

type OneCardOneAccountValidator struct {
	next *Validator
}

func (t *TwoElementValidator) Validate(params string) {
	fmt.Printf("process %v check two element validation\n", params)
	t.Next(params)
}

func (t *TwoElementValidator) SetNext(validator *Validator) {
	t.next = validator
}

func (t *TwoElementValidator) Next(params string)  {
	if t.next != nil {
		next := *(t.next)
		next.Validate(params)
	}
}

func (o *OneCardOneAccountValidator) Validate(params string) {
	fmt.Printf("process %v check one card one account validation\n", params)
	o.Next(params)
}

func (o *OneCardOneAccountValidator) SetNext(validator *Validator) {
	o.next = validator
}

func (o OneCardOneAccountValidator) Next(params string) {
	if o.next != nil {
		next := *(o.next)
		next.Validate(params)
	}
}



