package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
)

/*所与のロジック*/
type Dff struct {
	previousBit *Bit
}

func (gate *Dff) GetPreviousBit() *Bit {
	return gate.previousBit
}

func NewDff() *Dff {
	// TODO: registryのデフォ値はonだったような気がする(CPUの創り方を再読)
	return &Dff{previousBit: bit.OFF}
}

func (gate *Dff) Pass(in *Bit) *Bit {
	out := gate.previousBit
	gate.previousBit = in
	return out
}

func (gate *Dff) AsGate() bool { return true }
