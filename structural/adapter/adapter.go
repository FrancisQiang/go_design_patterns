package adapter

import (
	"fmt"
)

type SocketFor220V interface {
	Charge220V()
}

type SocketFor200V interface {
	Charge200V()
}

type RedBullSocket struct {
}

func (r RedBullSocket) Charge220V() {
	fmt.Println("charge for 220V")
}

type ChargerAdapter struct {
	socketFor220V SocketFor220V
}

func (c ChargerAdapter) Charge200V() {
	c.socketFor220V.Charge220V()
	fmt.Println("charge for 200V")
}



