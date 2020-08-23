package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestDemultiPlexer(t *testing.T) {
	gate := NewDeMultiPlexer()
	testcases := []struct {
		input     *Bit
		selector  *Bit
		expected1 *Bit
		expected2 *Bit
	}{
		{ON, ON, ON, OFF},
		{OFF, ON, OFF, OFF},
		{ON, OFF, OFF, ON},
		{OFF, OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actualout1, actualout2 := gate.Pass(testcase.input, testcase.selector)
		if actualout1 != testcase.expected1 || actualout2 != testcase.expected2 {
			t.Errorf("not match. caseno: %v, expected1: %+v, actual1: %+v, expected2: %+v, actual2: %+v",
				i, testcase.expected1, actualout1, testcase.expected2, actualout2)
		}
	}
}
