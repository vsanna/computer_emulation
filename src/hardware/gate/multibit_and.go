package gate

import . "computer_emulation/src/hardware/bit"

type MultibitAnd struct {
	and *And
}

func NewMultibitAnd() *MultibitAnd {
	return &MultibitAnd{and: NewAnd()}
}

func (gate *MultibitAnd) Pass(a *Bus, b *Bus) *Bus {
	out := NewBus(BusOption{})
	for idx, _ := range a.Bits {
		out.Bits[idx] = gate.and.Pass(a.Bits[idx], b.Bits[idx])
	}
	return out
}
func (gate *MultibitAnd) AsGate() bool { return true }
