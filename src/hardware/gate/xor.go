package gate

import . "computer_emulation/src/hardware/bit"

type Xor struct {
	and *And
	or  *Or
	not *Not
}

func NewXor() *Xor {
	return &Xor{and: NewAnd(), or: NewOr(), not: NewNot()}
}

func (gate *Xor) Pass(a *Bit, b *Bit) (out *Bit) {
	return gate.or.Pass(
		gate.and.Pass(a, gate.not.Pass(b)),
		gate.and.Pass(b, gate.not.Pass(a)),
	)
}
func (gate *Xor) AsGate() bool { return true }
