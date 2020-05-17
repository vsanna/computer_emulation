package gate

import . "computer_emulation/src/bit"

type MultibitNot struct {
	not *Not
}

func NewMultibitNot() *MultibitNot {
	return &MultibitNot{not: NewNot()}
}

func (gate *MultibitNot) Pass(in *Bus) *Bus {
	out := NewBus(BusOption{})
	for idx, _ := range in.Bits {
		out.Bits[idx] = gate.not.Pass(in.Bits[idx])
	}
	return out
}
func (gate *MultibitNot) AsGate() bool { return true }
