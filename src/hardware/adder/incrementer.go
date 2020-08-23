package adder

import (
	. "computer_emulation/src/hardware/bit"
)

type Incrementer struct {
	adder *Adder
}

func NewIncrementer() *Incrementer {
	return &Incrementer{adder: NewAdder()}
}

func (incrementer *Incrementer) Pass(a *Bus) *Bus {
	oneBus := NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}})
	return incrementer.adder.Pass(a, oneBus)
}
