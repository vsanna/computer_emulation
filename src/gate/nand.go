package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
)

/*所与のロジック*/
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
