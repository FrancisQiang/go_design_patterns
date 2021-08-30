package iterator

type Element struct {
	value string
}

type Iterator interface {
	HasNext() bool
	Next() *Element
}

type ElementCollection struct {
	elements []*Element
}

func (ec *ElementCollection) Iterator() Iterator {
	return &ElementIterator{
		elements: ec.elements,
	}
}

type ElementIterator struct {
	index    int
	elements []*Element
}

func (e *ElementIterator) HasNext() bool {
	if e.index < len(e.elements) {
		return true
	}
	return false
}

func (e *ElementIterator) Next() *Element {
	if e.HasNext() {
		element := e.elements[e.index]
		e.index++
		return element
	}
	return nil
}

