package adder

import (
	. "computer_emulation/src/hardware/bit"
)

type Adder struct {
	full_adder *FullAdder
}

func NewAdder() *Adder {
	return &Adder{full_adder: NewFullAdder()}
}

func (adder *Adder) Pass(a *Bus, b *Bus) *Bus {
	bus := NewBus(BusOption{})
	sum := OFF
	cb := OFF
	for idx, _ := range a.Bits {
		pos := BUS_WIDTH - 1 - idx
		sum, cb = adder.full_adder.Pass(a.Bits[pos], b.Bits[pos], cb)
		bus.Bits[pos] = sum
	}
	return bus
}
