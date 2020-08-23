package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestMultibitOr(t *testing.T) {
	gate := NewMultibitOr()
	testcases := []struct {
		input1   *Bus
		input2   *Bus
		expected *Bus
	}{
		{
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}),
			NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}}),
		},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input1, testcase.input2)
		if !actual.Equals(testcase.expected) {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
