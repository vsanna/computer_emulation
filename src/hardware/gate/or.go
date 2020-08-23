package gate

import . "computer_emulation/src/hardware/bit"

type Or struct {
	nand *Nand
	not  *Not
}

func NewOr() *Or {
	return &Or{nand: NewNand(), not: NewNot()}
}

func (gate *Or) Pass(a *Bit, b *Bit) (out *Bit) {
	return gate.nand.Pass(gate.not.Pass(a), gate.not.Pass(b))
}
func (gate *Or) AsGate() bool { return true }
