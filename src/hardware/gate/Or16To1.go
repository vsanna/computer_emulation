package gate

import (
	. "computer_emulation/src/hardware/bit"
)

type Or16To1 struct {
	or *Or
}

func NewOr16To1() *Or16To1 {
	return &Or16To1{or: NewOr()}
}

func (gate *Or16To1) Pass(a *Bus) (out *Bit) {
	aggregate := OFF
	for _, bit := range a.Bits {
		aggregate = gate.or.Pass(aggregate, bit)
	}
	return aggregate
}
func (gate *Or16To1) AsGate() bool { return true }
