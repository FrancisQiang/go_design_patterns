package command

import (
	"fmt"
)

type Command interface {
	Execute()
}

type CopyCommand struct {

}

func (c CopyCommand) Execute() {
	fmt.Println("copy")
}

type DeleteCommand struct {

}

func (d DeleteCommand) Execute() {
	fmt.Println("delete")
}

type Button struct {
	command Command
}

func (b Button) SetCommand(command Command) {
	b.command = command
}

func (b Button) onPressed() {
	b.command.Execute()
}

