package gate

import . "computer_emulation/src/bit"

type MultibitOr struct {
	or *Or
}

func NewMultibitOr() *MultibitOr {
	return &MultibitOr{or: NewOr()}
}

func (gate *MultibitOr) Pass(a *Bus, b *Bus) *Bus {
	out := NewBus(BusOption{})
	for idx, _ := range a.Bits {
		out.Bits[idx] = gate.or.Pass(a.Bits[idx], b.Bits[idx])
	}
	return out
}
func (gate *MultibitOr) AsGate() bool { return true }
