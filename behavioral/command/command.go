package command

import (
	"fmt"
)
// 使命令和调用者解耦
// 抽象一个命令类，有很多子类去实现
// 调用者类持有命令类基类
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

