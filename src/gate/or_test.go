package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestOr(t *testing.T) {
	gate := NewOr()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		expected *Bit
	}{
		{bit.ON, bit.ON, bit.ON},
		{bit.ON, bit.OFF, bit.ON},
		{bit.OFF, bit.ON, bit.ON},
		{bit.OFF, bit.OFF, bit.OFF},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input1, testcase.input2)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
