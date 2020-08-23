package memory

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestWord(t *testing.T) {
	testcases := []struct {
		load     *Bit
		input    *Bus
		expected *Bus
	}{
		{
			ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
		},
		{
			ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
		},
		{
			OFF,
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
		},
		{
			ON,
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
		},
		{
			ON,
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
			NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
		},
	}
	word := NewWord()

	for i, testcase := range testcases {
		actual := word.Pass(testcase.input, testcase.load)
		if !actual.Equals(testcase.expected) {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
