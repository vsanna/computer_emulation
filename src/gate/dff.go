package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
)

// dff is provided. this is one of the two physical basis.
type Dff struct {
	previousBit *Bit
}

func (gate *Dff) GetPreviousBit() *Bit {
	return gate.previousBit
}

func NewDff() *Dff {
	return &Dff{previousBit: bit.OFF}
}

func (gate *Dff) Pass(in *Bit) *Bit {
	out := gate.previousBit
	gate.previousBit = in
	return out
}

func (gate *Dff) AsGate() bool { return true }
