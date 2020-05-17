package adder

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
)

type FullAdder struct {
	half_adder *HalfAdder
	or         *gate.Or
}

func NewFullAdder() *FullAdder {
	return &FullAdder{half_adder: NewHalfAdder(), or: gate.NewOr()}
}

func (adder *FullAdder) Pass(a *Bit, b *Bit, c *Bit) (sum *Bit, carryover *Bit) {
	sum1, cb1 := adder.half_adder.Pass(a, b)
	sum2, cb2 := adder.half_adder.Pass(sum1, c)
	carryover = adder.or.Pass(cb1, cb2)
	return sum2, carryover
}
