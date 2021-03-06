package memory

import (
	. "computer_emulation/src/hardware/bit"
	"computer_emulation/src/hardware/gate"
)

type RegisterCell struct {
	dff          *gate.Dff
	multi_prexer *gate.MultiPlexer
}

func NewRegisterCell() *RegisterCell {
	return &RegisterCell{dff: gate.NewDff(), multi_prexer: gate.NewMultiPlexer()}
}

// load:1 is write mode
func (register *RegisterCell) Pass(in *Bit, load *Bit) *Bit {
	out := register.dff.GetPreviousBit()
	newIn := register.multi_prexer.Pass(in, out, load)
	register.dff.Pass(newIn)
	return out
}

func (register *RegisterCell) String() string {
	return register.dff.GetPreviousBit().String()
}
