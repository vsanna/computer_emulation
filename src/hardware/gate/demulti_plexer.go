package gate

import . "computer_emulation/src/hardware/bit"

type DeMultiPlexer struct {
	not *Not
	and *And
}

func NewDeMultiPlexer() *DeMultiPlexer {
	return &DeMultiPlexer{not: NewNot(), and: NewAnd()}
}

// selector: 1 -> out1, selector: 2 -> out2
func (gate *DeMultiPlexer) Pass(in *Bit, selector *Bit) (out1 *Bit, out2 *Bit) {
	out1 = gate.and.Pass(selector, in)
	out2 = gate.and.Pass(gate.not.Pass(selector), in)
	return
}
func (gate *DeMultiPlexer) AsGate() bool { return true }
