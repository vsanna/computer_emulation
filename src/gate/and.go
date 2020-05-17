package gate

import . "computer_emulation/src/bit"

type And struct {
	nand *Nand
	not  *Not
}

func NewAnd() *And {
	return &And{nand: NewNand(), not: NewNot()}
}

func (gate *And) Pass(a *Bit, b *Bit) (out *Bit) {
	return gate.not.Pass(gate.nand.Pass(a, b))
}
func (gate *And) AsGate() bool { return true }
