package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestNot(t *testing.T) {
	gate := NewNot()
	testcases := []struct {
		input    *Bit
		expected *Bit
	}{
		{bit.ON, bit.OFF},
		{bit.OFF, bit.ON},
	}

	for i, testcase := range testcases {
		actual := gate.Pass(testcase.input)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
