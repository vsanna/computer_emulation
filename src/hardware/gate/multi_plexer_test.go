package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestMultiPlexer(t *testing.T) {
	gate := NewMultiPlexer()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		selector *Bit
		expected *Bit
	}{
		{ON, ON, ON, ON},
		{OFF, ON, ON, OFF},
		{ON, OFF, ON, ON},
		{OFF, OFF, ON, OFF},
		{ON, ON, OFF, ON},
		{ON, OFF, OFF, OFF},
		{OFF, ON, OFF, ON},
		{OFF, OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input1, testcase.input2, testcase.selector)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
