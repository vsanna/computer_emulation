package gate

import (
	. "computer_emulation/src/hardware/bit"
)

// nand is provided. this is one of the two physical basis.
type Nand struct{}

func NewNand() *Nand {
	return &Nand{}
}

func (dff *Nand) Pass(a *Bit, b *Bit) (out *Bit) {
	if a == ON && b == ON {
		return OFF
	}
	return ON
}
func (dff *Nand) AsGate() bool { return true }
