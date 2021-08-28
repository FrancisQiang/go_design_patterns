package template

import (
	"fmt"
)

type Worker interface {
	GetUp()
	Work()
	Sleep()
}

type Daily struct {
	worker Worker
}

func (daily *Daily) Live() {
	daily.worker.GetUp()
	daily.worker.Work()
	daily.worker.Sleep()
}

type Coder struct {

}

func (c Coder) GetUp() {
	fmt.Println("get up")
}

func (c Coder) Work() {
	fmt.Println("coding")
}

func (c Coder) Sleep() {
	fmt.Println("sleeping")
}

