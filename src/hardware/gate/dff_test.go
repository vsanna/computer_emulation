package gate

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestDff(t *testing.T) {
	testcases := []struct {
		input1    *Bit
		input2    *Bit
		expected0 *Bit
		expected1 *Bit
		expected2 *Bit
	}{
		{ON, OFF, OFF, ON, OFF},
		{ON, ON, OFF, ON, ON},
		{OFF, ON, OFF, OFF, ON},
		{OFF, OFF, OFF, OFF, OFF},
	}

	for i, testcase := range testcases {
		gate := NewDff()
		actual0 := gate.Pass(testcase.input1)
		if actual0 != testcase.expected0 {
			t.Errorf("not match. caseno: %v, expected0: %+v, actual0: %+v", i, testcase.expected0, actual0)
		}

		actual1 := gate.Pass(testcase.input2)
		if actual1 != testcase.expected1 {
			t.Errorf("not match. caseno: %v, expected1: %+v, actual1: %+v", i, testcase.expected1, actual1)

		}

		actual2 := gate.Pass(ON)
		if actual2 != testcase.expected2 {
			t.Errorf("not match. caseno: %v, expected2: %+v, actual2: %+v", i, testcase.expected2, actual2)
		}
	}
}
