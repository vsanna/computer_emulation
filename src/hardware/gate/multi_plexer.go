package gate

import . "computer_emulation/src/hardware/bit"

type MultiPlexer struct {
	not *Not
	or  *Or
	and *And
}

// 2 bits -> 1 bit
func NewMultiPlexer() *MultiPlexer {
	return &MultiPlexer{not: NewNot(), or: NewOr(), and: NewAnd()}
}

// selector: 1 -> a, selector: 0 -> b
func (gate *MultiPlexer) Pass(a *Bit, b *Bit, selector *Bit) (out *Bit) {
	return gate.or.Pass(
		gate.and.Pass(selector, a),
		gate.and.Pass(gate.not.Pass(selector), b),
	)
}
func (gate *MultiPlexer) AsGate() bool { return true }
