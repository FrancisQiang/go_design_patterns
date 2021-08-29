package observer

import (
	"fmt"
)

type IObserver interface {
	Update()
}

type AObserver struct {
}

type BObserver struct {
}

type CObserver struct {
}

func (A AObserver) Update() {
	fmt.Println("A update")
}

func (B BObserver) Update() {
	fmt.Println("B update")
}

func (C CObserver) Update() {
	fmt.Println("C update")
}

type Observer struct {
	subject   *Subject
	iObserver IObserver
}

func NewObserver(observer IObserver, subject *Subject) Observer {
	o := Observer{
		iObserver: observer,
		subject:   subject,
	}
	subject.Attach(o)
	return o
}

type Subject struct {
	observers []Observer
}

func NewSubject() *Subject {
	return &Subject{
		observers: []Observer{},
	}
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) SetState() {
	fmt.Println("Begin to set state")
	s.notifyAllObservers()
}

func (s *Subject) notifyAllObservers() {
	for _, o := range s.observers {
		o.iObserver.Update()
	}
}
