package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestNewMultibitTo1MultiPlexer(t *testing.T) {
	gate := NewMultibitTo1MultiPlexer()
	testcases := []struct {
		input    *Bus
		s1       *Bit
		s2       *Bit
		s3       *Bit
		s4       *Bit
		expected *Bit
	}{
		{
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}),
			OFF, OFF, OFF, OFF,
			ON,
		},
		{
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}}),
			OFF, OFF, OFF, ON,
			ON,
		},
		{
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}}),
			OFF, OFF, ON, OFF,
			ON,
		},
		{
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}}),
			OFF, OFF, ON, ON,
			ON,
		},
		{
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}}),
			OFF, ON, OFF, OFF,
			ON,
		},
		{
			bit.NewBus(BusOption{Bits: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
			ON, ON, ON, ON,
			ON,
		},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input, testcase.s1, testcase.s2, testcase.s3, testcase.s4)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
