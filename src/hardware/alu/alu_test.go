package alu

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestAlu(t *testing.T) {
	alu := NewAlu()
	testcases := []struct {
		a            *Bus
		b            *Bus
		zerox        *Bit
		negatex      *Bit
		zeroy        *Bit
		negatey      *Bit
		f            *Bit
		negateoutput *Bit

		expectedOut              *Bus
		expectedOutputIsZero     *Bit
		expectedOutputIzNegative *Bit
	}{
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			OFF,
			ON,
			OFF,
			ON,
			OFF,

			// 0
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			ON,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			ON,
			ON,
			ON,
			ON,

			// 1
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			ON,
			OFF,
			ON,
			OFF,

			// -1
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			OFF,
			ON,
			ON,
			OFF,
			OFF,

			// x
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			OFF,
			OFF,
			OFF,
			OFF,

			// y
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			OFF,
			ON,
			ON,
			OFF,
			ON,

			// "!x"
			NewBus(BusOption{Bits: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			ON,
			ON,
			OFF,
			OFF,
			OFF,
			ON,

			// "!y"
			NewBus(BusOption{Bits: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			OFF,
			OFF,
			ON,
			ON,
			ON,
			ON,

			// -x
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			OFF,
			OFF,
			ON,
			ON,

			// -x
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			ON,
			ON,
			ON,
			ON,
			ON,

			// x+1
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			OFF,
			ON,
			ON,
			ON,

			// y+1
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			OFF,
			ON,
			ON,
			ON,
			OFF,

			// x-1
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			ON,
			ON,
			OFF,
			OFF,
			ON,
			OFF,

			// y-1
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			OFF,
			OFF,
			OFF,
			ON,
			OFF,

			// x+y
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			OFF,
			ON,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}}),
			OFF,
			ON,
			OFF,
			OFF,
			ON,
			ON,

			// x-y
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}}),
			OFF,
			OFF,
			OFF,
			ON,
			ON,
			ON,

			// y-x
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}}),
			OFF,
			OFF,
			OFF,
			OFF,
			OFF,
			OFF,

			// x & y
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}}),
			OFF,
			OFF,
		},
		{
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}}),
			OFF,
			ON,
			OFF,
			ON,
			OFF,
			ON,

			// x | y
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}}),
			OFF,
			OFF,
		},
	}

	for i, testcase := range testcases {
		actualOutput, actualOutputIsZero, actualOutputIsNegative := alu.Pass(
			testcase.a,
			testcase.b,
			testcase.zerox,
			testcase.negatex,
			testcase.zeroy,
			testcase.negatey,
			testcase.f,
			testcase.negateoutput,
		)
		if !actualOutput.Equals(testcase.expectedOut) ||
			actualOutputIsZero != testcase.expectedOutputIsZero ||
			actualOutputIsNegative != testcase.expectedOutputIzNegative {
			t.Errorf("not match. caseno: %v\n\t"+
				"expectedOut: %+v, actualOutput: %+v\n\t"+
				"expectedOutputIsZero: %+v, actualOutputIsZero: %+v\n\t"+
				"expectedOutputIzNegative: %+v, actualOutputIsNegative: %+v\n\t",
				i, testcase.expectedOut, actualOutput,
				testcase.expectedOutputIsZero, actualOutputIsZero,
				testcase.expectedOutputIzNegative, actualOutputIsNegative)
		}
	}
}
