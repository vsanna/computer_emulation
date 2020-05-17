package alu

import (
	"computer_emulation/src/adder"
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
)

type Jump struct {
	multibit_multi_plexer *gate.MultibitMultiPlexer
	or                    *gate.Or
	and                   *gate.And
	not                   *gate.Not
	incrementer           *adder.Incrementer
}

func NewJump() *Jump {
	return &Jump{
		multibit_multi_plexer: gate.NewMultibitMultiPlexer(),
		or:                    gate.NewOr(),
		and:                   gate.NewAnd(),
		not:                   gate.NewNot(),
		incrementer:           adder.NewIncrementer(),
	}
}

func (jump *Jump) Pass(
	currentA *Bus,
	currentPC *Bus,
	outputIsZero *Bit,
	outputIsNegative *Bit,
	jumpBus *Bus,
	reset *Bit,
) *Bus {
	jgtBit := jump.and.Pass(
		jump.not.Pass(jumpBus.Bits[0]),
		jump.and.Pass(
			jump.not.Pass(jumpBus.Bits[1]),
			jump.and.Pass(
				jumpBus.Bits[2],
				jump.and.Pass(
					jump.not.Pass(outputIsNegative),
					jump.not.Pass(outputIsZero)))))

	jeqBit := jump.and.Pass(
		jump.not.Pass(jumpBus.Bits[0]),
		jump.and.Pass(
			jumpBus.Bits[1],
			jump.and.Pass(
				jump.not.Pass(jumpBus.Bits[2]),
				outputIsZero)))

	jgeBit := jump.and.Pass(
		jump.not.Pass(jumpBus.Bits[0]),
		jump.and.Pass(
			jumpBus.Bits[1],
			jump.and.Pass(
				jumpBus.Bits[2],
				jump.or.Pass(
					jump.not.Pass(outputIsNegative),
					outputIsZero))))

	jltBit := jump.and.Pass(
		jumpBus.Bits[0],
		jump.and.Pass(
			jump.not.Pass(jumpBus.Bits[1]),
			jump.and.Pass(
				jump.not.Pass(jumpBus.Bits[2]),
				outputIsNegative)))

	jneBit := jump.and.Pass(
		jumpBus.Bits[0],
		jump.and.Pass(
			jump.not.Pass(jumpBus.Bits[1]),
			jump.and.Pass(
				jumpBus.Bits[2],
				jump.not.Pass(outputIsZero))))

	jleBit := jump.and.Pass(
		jumpBus.Bits[0],
		jump.and.Pass(
			jumpBus.Bits[1],
			jump.and.Pass(
				jump.not.Pass(jumpBus.Bits[2]),
				jump.or.Pass(
					outputIsNegative,
					outputIsZero))))

	alwaysJumpBit := jump.and.Pass(jumpBus.Bits[0], jump.and.Pass(jumpBus.Bits[1], jumpBus.Bits[2]))

	shuoldJumpBit := jump.or.Pass(jgtBit,
		jump.or.Pass(jeqBit,
			jump.or.Pass(jgeBit,
				jump.or.Pass(jltBit,
					jump.or.Pass(jneBit,
						jump.or.Pass(jleBit, alwaysJumpBit))))))

	return jump.multibit_multi_plexer.Pass(
		NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
		jump.multibit_multi_plexer.Pass(
			currentA,
			jump.incrementer.Pass(currentPC),
			shuoldJumpBit,
		),
		reset,
	)
}
