package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestMultibitNot(t *testing.T) {
	gate := NewMultibitNot()
	testcases := []struct {
		input    *Bus
		expected *Bus
	}{
		{
			bit.NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}}),
			bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}}),
		},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input)
		if !actual.Equals(testcase.expected) {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
