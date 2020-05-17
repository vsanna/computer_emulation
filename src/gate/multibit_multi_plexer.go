package gate

import . "computer_emulation/src/bit"

// 2 bus -> 1 bus
type MultibitMultiPlexer struct {
	multi_plexer *MultiPlexer
}

func NewMultibitMultiPlexer() *MultibitMultiPlexer {
	return &MultibitMultiPlexer{multi_plexer: NewMultiPlexer()}
}

// selector: 1 -> a, selector: 0 -> b
func (gate *MultibitMultiPlexer) Pass(a *Bus, b *Bus, selector *Bit) (out *Bus) {
	bus := NewBus(BusOption{})
	for idx, _ := range a.Bits {
		bus.Bits[idx] = gate.multi_plexer.Pass(a.Bits[idx], b.Bits[idx], selector)
	}
	return bus
}
func (gate *MultibitMultiPlexer) AsGate() bool { return true }
