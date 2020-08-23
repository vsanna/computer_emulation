package alu

import (
	. "computer_emulation/src/hardware/bit"
)

type Dest struct {
}

func NewDest() *Dest {
	return &Dest{}
}

func (dest *Dest) Pass(in *Bus) (loadA *Bit, loadD *Bit, loadM *Bit) {
	return in.Bits[0], in.Bits[1], in.Bits[2]
}
