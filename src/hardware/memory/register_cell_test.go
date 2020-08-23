package memory

import (
	. "computer_emulation/src/hardware/bit"
	"testing"
)

func TestRegisterCell(t *testing.T) {
	testcases := []struct {
		input    *Bit
		load     *Bit
		expected *Bit
	}{
		{ON, ON, OFF},
		{OFF, ON, ON},   // 0でONに上書き
		{ON, OFF, OFF},  // 1でOFFに上書き
		{ON, OFF, OFF},  // 2で上書きなし
		{OFF, OFF, OFF}, // 3で上書きなし
	}
	register := NewRegisterCell()

	for i, testcase := range testcases {
		actual := register.Pass(testcase.input, testcase.load)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
