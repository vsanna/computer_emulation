package alu

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
)

type Comp struct {
	multibit_multi_plexer *gate.MultibitMultiPlexer
}

func NewComp() *Comp {
	return &Comp{multibit_multi_plexer: gate.NewMultibitMultiPlexer()}
}

// inのうち、下7桁のみ利用する
// multibit_multi_plexerを使い、0, D, A, Mから2つ選ぶ
// multibit_multi_plexer(0, d_reg, a_reg, memory[a_reg])
// またALU向け指示コードを返す
func (dest *Comp) Pass(
	a_reg_bus *Bus,
	d_reg_bus *Bus,
	m_bus *Bus,
	in *Bus,
) (x *Bus, y *Bus, operationCode *Bus) {
	// Dは確定利用
	x = d_reg_bus
	// A or Mは先頭1bit次第
	y = dest.multibit_multi_plexer.Pass(
		a_reg_bus,
		m_bus,
		in.Bits[0],
	)
	operationCode = NewBus(BusOption{
		Bits: []int{in.Bits[1].GetVal(), in.Bits[2].GetVal(), in.Bits[3].GetVal(), in.Bits[4].GetVal(), in.Bits[5].GetVal(), in.Bits[6].GetVal(), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	return x, y, operationCode
}
