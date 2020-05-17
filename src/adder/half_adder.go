package adder

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
)

type HalfAdder struct {
	xor *gate.Xor
	and *gate.And
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{xor: gate.NewXor(), and: gate.NewAnd()}
}

func (adder *HalfAdder) Pass(a *Bit, b *Bit) (sum *Bit, carryover *Bit) {
	sum = adder.xor.Pass(a, b)
	carryover = adder.and.Pass(a, b)
	return sum, carryover
}
