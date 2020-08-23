package alu

import (
	"computer_emulation/src/hardware/adder"
	. "computer_emulation/src/hardware/bit"
	"computer_emulation/src/hardware/gate"
)

type Alu struct {
	multi_plexer               *gate.MultibitMultiPlexer
	adder                      *adder.Adder
	and                        *gate.MultibitAnd
	not                        *gate.MultibitNot
	not_single_bit             *gate.Not
	or16To1                    *gate.Or16To1
	multibit_to_1_multi_plexer *gate.MultibitTo1MultiPlexer
}

func NewAlu() *Alu {
	return &Alu{
		multi_plexer:               gate.NewMultibitMultiPlexer(),
		not:                        gate.NewMultibitNot(),
		not_single_bit:             gate.NewNot(),
		or16To1:                    gate.NewOr16To1(),
		and:                        gate.NewMultibitAnd(),
		adder:                      adder.NewAdder(),
		multibit_to_1_multi_plexer: gate.NewMultibitTo1MultiPlexer(),
	}
}

// ALU has 2 16bit buses and 6 flags as input
// theoritically, it can do calculation in 2**6 = 64 ways,
// but it supports only 18ways as this hardware's specification.
func (alu *Alu) Pass(
	a *Bus, b *Bus,
	zerox *Bit, negatex *Bit,
	zeroy *Bit, negatey *Bit,
	f *Bit,
	negateoutput *Bit) (out *Bus, outputIsZero *Bit, outputIsNegative *Bit) {
	zerofyA := alu.multi_plexer.Pass(
		NewBus(BusOption{}),
		a,
		zerox,
	)
	negatedA := alu.multi_plexer.Pass(
		alu.not.Pass(zerofyA),
		zerofyA,
		negatex,
	)

	zerofyB := alu.multi_plexer.Pass(
		NewBus(BusOption{}),
		b,
		zeroy,
	)
	negatedB := alu.multi_plexer.Pass(
		alu.not.Pass(zerofyB),
		zerofyB,
		negatey,
	)

	functionApplied := alu.multi_plexer.Pass(
		alu.adder.Pass(negatedA, negatedB),
		alu.and.Pass(negatedA, negatedB),
		f,
	)

	out = alu.multi_plexer.Pass(
		alu.not.Pass(functionApplied),
		functionApplied,
		negateoutput,
	)

	// outputIsZero
	// consolidate 16bit bus into 1bit
	outputIsZero = alu.not_single_bit.Pass(alu.or16To1.Pass(out))

	// outputIsNegative
	outputIsNegative = alu.multibit_to_1_multi_plexer.Pass(out, ON, ON, ON, ON)

	return out, outputIsZero, outputIsNegative
}
