package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestNand(t *testing.T) {
	nand := NewNand()
	testcases := []struct {
		input1   *Bit
		input2   *Bit
		expected *Bit
	}{
		{ON, ON, OFF},
		{ON, OFF, ON},
		{OFF, ON, ON},
		{OFF, OFF, ON},
	}

	for i, testcase := range testcases {
		actual := nand.Pass(testcase.input1, testcase.input2)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
