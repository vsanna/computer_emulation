package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestAnd(t *testing.T) {
	gate := NewAnd()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		expected *Bit
	}{
		{ON, ON, ON},
		{ON, OFF, OFF},
		{OFF, ON, OFF},
		{OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input1, testcase.input2)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
