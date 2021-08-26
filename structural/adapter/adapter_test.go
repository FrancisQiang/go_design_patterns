package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	socket := &RedBullSocket{}
	socket.Charge220V()
	adapterSocket := &ChargerAdapter{
		RedBullSocket: *socket,
	}
	adapterSocket.Charge200V()
}
