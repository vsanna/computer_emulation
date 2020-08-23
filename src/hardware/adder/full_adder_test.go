package adder

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestFullAdder(t *testing.T) {
	adder := NewFullAdder()
	testcases := []struct {
		input1       *Bit
		input2       *Bit
		input3       *Bit
		expected_sum *Bit
		expected_cb  *Bit
	}{
		{ON, ON, ON, ON, ON},

		{OFF, ON, ON, OFF, ON},
		{ON, OFF, ON, OFF, ON},
		{ON, ON, OFF, OFF, ON},

		{OFF, OFF, ON, ON, OFF},
		{ON, OFF, OFF, ON, OFF},
		{OFF, ON, OFF, ON, OFF},

		{OFF, OFF, OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actual_sum, actual_cb := adder.Pass(testcase.input1, testcase.input2, testcase.input3)
		if actual_sum != testcase.expected_sum || actual_cb != testcase.expected_cb {
			t.Errorf("not match. caseno: %v, expected_sum: %+v, actual_sum: %+v, expected_cb: %+v, actual_cb: %+v",
				i, testcase.expected_sum, actual_sum, testcase.expected_cb, actual_cb)
		}
	}
}
