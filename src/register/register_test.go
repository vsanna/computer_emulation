package register

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"math"
	"testing"
)

func TestRegister(t *testing.T) {
	testcases := []struct {
		load     *Bit
		input    *Bus
		expected *Bus
	}{
		{
			bit.ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
		},
		{
			bit.ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
		},
		{
			bit.OFF,
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
		},
		{
			bit.ON,
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
		},
		{
			bit.ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
		},
	}
	register := NewRegister("TEST")

	for i, testcase := range testcases {
		actual := register.Pass(testcase.input, testcase.load)
		if !actual.Equals(testcase.expected) {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}

func TestRegisterToInt(t *testing.T) {
	reg1 := NewRegister("test1")
	reg1.Pass(NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}), ON)

	reg2 := NewRegister("test2")
	reg2.Pass(NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}}), ON)

	reg3 := NewRegister("test3")
	reg3.Pass(NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}), ON)

	reg4 := NewRegister("test4")
	reg4.Pass(NewBus(BusOption{Bits: []int{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}}), ON)

	testcases := []struct {
		input    *Register
		expected int
	}{
		{reg1, 0},
		{reg2, 3},
		{reg3, 65535},
		{reg4, (int)(1 + 2 + math.Pow(2, 11) + math.Pow(2, 12))},
	}

	for i, testcase := range testcases {
		if testcase.input.ToInt() != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, testcase.input.ToInt())
		}
	}
}
