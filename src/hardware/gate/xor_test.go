package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestXor(t *testing.T) {
	gate := NewXor()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		expected *Bit
	}{
		{ON, ON, OFF},
		{ON, OFF, ON},
		{OFF, ON, ON},
		{OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input1, testcase.input2)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
