package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	s := NewSubject()
	NewObserver(AObserver{}, s)
	NewObserver(BObserver{}, s)
	NewObserver(CObserver{}, s)
	s.SetState()

}
