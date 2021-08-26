package proxy

import (
	"fmt"
)

type Task interface {
	Execute()
}

type BuyTicketTask struct {

}

func (b BuyTicketTask) Execute() {
	fmt.Println("buy a ticket")
}

type BuyTicketAgent struct {
	task Task
}

func (b BuyTicketAgent) Execute() {
	fmt.Println("agent do")
	b.task.Execute()
}

