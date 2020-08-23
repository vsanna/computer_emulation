package adder

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	adder := NewHalfAdder()
	testcases := []struct {
		input1       *Bit
		input2       *Bit
		expected_sum *Bit
		expected_cb  *Bit
	}{
		{ON, ON, OFF, ON},
		{ON, OFF, ON, OFF},
		{OFF, ON, ON, OFF},
		{OFF, OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		actual_sum, actual_cb := adder.Pass(testcase.input1, testcase.input2)
		if actual_sum != testcase.expected_sum || actual_cb != testcase.expected_cb {
			t.Errorf("not match. caseno: %v, expected_sum: %+v, actual_sum: %+v, expected_cb: %+v, actual_cb: %+v",
				i, testcase.expected_sum, actual_sum, testcase.expected_cb, actual_cb)
		}
	}
}
