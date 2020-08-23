package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
)

// nand is provided. this is one of the two physical basis.
type Nand struct{}

func NewNand() *Nand {
	return &Nand{}
}

func (dff *Nand) Pass(a *Bit, b *Bit) (out *Bit) {
	if a == bit.ON && b == bit.ON {
		return bit.OFF
	}
	return bit.ON
}
func (dff *Nand) AsGate() bool { return true }
