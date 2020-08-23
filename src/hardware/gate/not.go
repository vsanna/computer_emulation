package gate

import . "computer_emulation/src/hardware/bit"

type Not struct {
	nand *Nand
}

func NewNot() *Not { return &Not{nand: NewNand()} }

func (gate *Not) Pass(in *Bit) (out *Bit) {
	return gate.nand.Pass(in, in)
}
func (gate *Not) AsGate() bool { return true }
