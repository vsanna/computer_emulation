package alu

import (
	. "computer_emulation/src/hardware/bit"
	"computer_emulation/src/hardware/gate"
)

type Comp struct {
	multibit_multi_plexer *gate.MultibitMultiPlexer
}

func NewComp() *Comp {
	return &Comp{multibit_multi_plexer: gate.NewMultibitMultiPlexer()}
}

// use only heading 7bit of in bus.(since a comp part is composed of 7bit of an instruction)
// according to the comp part, choose 2 input source from D + (A or M)
// and retrieve operationCode(6bits) which is based on ALU spec.
func (dest *Comp) Pass(
	a_reg_bus *Bus,
	d_reg_bus *Bus,
	m_bus *Bus,
	in *Bus,
) (x *Bus, y *Bus, operationCode *Bus) {
	// Dはソースの一方として確定利用(その後ALUでzerofyされる可能性はある)
	x = d_reg_bus

	// 他方の入力ソースはA or M. 先頭1bit次第
	y = dest.multibit_multi_plexer.Pass(
		m_bus,
		a_reg_bus,
		in.Bits[0],
	)

	operationCode = NewBus(BusOption{
		Bits: []int{
			in.Bits[1].GetVal(),
			in.Bits[2].GetVal(),
			in.Bits[3].GetVal(),
			in.Bits[4].GetVal(),
			in.Bits[5].GetVal(),
			in.Bits[6].GetVal(),
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	return x, y, operationCode
}
