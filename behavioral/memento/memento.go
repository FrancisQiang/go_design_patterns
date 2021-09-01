package memento

type Memento struct {
	state string
}

func NewMemento(st string) *Memento {
	return &Memento{
		state: st,
	}
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	state string
}

func NewOriginator(st string) *Originator {
	return &Originator{
		state: st,
	}
}

func (o *Originator) SetState(st string) {
	o.state = st
}


func (o *Originator) GetState() string {
	return o.state
}


func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.state)
}

func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
}


type CareTaker struct {
	MementoList map[int]*Memento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{
		MementoList: make(map[int]*Memento),
	}
}

func (ct *CareTaker) Add(index int, memento *Memento) {
	ct.MementoList[index] = memento
}

func (ct *CareTaker) Get(index int) *Memento {
	return ct.MementoList[index]
}