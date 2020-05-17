package memory

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"testing"
)

func TestRegisterCell(t *testing.T) {
	testcases := []struct {
		input    *Bit
		load     *Bit
		expected *Bit
	}{
		{bit.ON, bit.ON, bit.OFF},
		{bit.OFF, bit.ON, bit.ON},   // 0でONに上書き
		{bit.ON, bit.OFF, bit.OFF},  // 1でOFFに上書き
		{bit.ON, bit.OFF, bit.OFF},  // 2で上書きなし
		{bit.OFF, bit.OFF, bit.OFF}, // 3で上書きなし
	}
	register := NewRegisterCell()

	for i, testcase := range testcases {
		actual := register.Pass(testcase.input, testcase.load)
		if actual != testcase.expected {
			t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
		}
	}
}
