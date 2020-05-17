package alu

import (
	"computer_emulation/src/adder"
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
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

// ALUは16bitの入力2系統に、6つのflagを入力として受け取る
// 理論的には2**6 = 64通りの演算を行えるが、サポートするのはそのうち18種類とする(このhardwareでは。)
func (alu *Alu) Pass(
	a *Bus, b *Bus,
	zerox *Bit, negatex *Bit,
	zeroy *Bit, negatey *Bit,
	f *Bit,
	negateoutput *Bit) (out *Bus, outputIsZero *Bit, outputIsNegative *Bit) {
	zerofyA := alu.multi_plexer.Pass(
		bit.NewBus(BusOption{}),
		a,
		zerox,
	)
	negatedA := alu.multi_plexer.Pass(
		alu.not.Pass(zerofyA),
		zerofyA,
		negatex,
	)

	zerofyB := alu.multi_plexer.Pass(
		bit.NewBus(BusOption{}),
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
	// Orを16(15)使い1Bitの出力にまとめ上げる
	outputIsZero = alu.not_single_bit.Pass(alu.or16To1.Pass(out))

	// outputIsNegative
	outputIsNegative = alu.multibit_to_1_multi_plexer.Pass(out, ON, ON, ON, ON)

	return out, outputIsZero, outputIsNegative
}
