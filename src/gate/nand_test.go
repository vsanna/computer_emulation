package gate

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestNand(t *testing.T) {
	nand := NewNand()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		expected *Bit
	}{
		{bit.ON, bit.ON, bit.OFF},
		{bit.ON, bit.OFF, bit.ON},
		{bit.OFF, bit.ON, bit.ON},
		{bit.OFF, bit.OFF, bit.ON},
	}

	for i, testcase := range testcases {
		actual := nand.Pass(testcase.input1, testcase.input2)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
