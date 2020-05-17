package memory

import (
	. "computer_emulation/src/bit"
	"testing"
)

// TODO: add test case
func TestMemory(t *testing.T) {
	// testcases := []struct {
	// 	load     *Bit
	// 	input    *Bus
	// 	expected *Bus
	// }{
	// 	{
	// 		bit.ON,
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
	// 		NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
	// 	},
	// 	{
	// 		bit.ON,
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
	// 	},
	// 	{
	// 		bit.OFF,
	// 		NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
	// 	},
	// 	{
	// 		bit.ON,
	// 		NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}}),
	// 	},
	// 	{
	// 		bit.ON,
	// 		NewBus(BusOption{Bits: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
	// 		NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}),
	// 	},
	// }
	// word := NewWord()

	// for i, testcase := range testcases {
	// 	actual := word.Pass(testcase.input, testcase.load)
	// 	if !actual.Equals(testcase.expected) {
	// 		t.Errorf("not match. caseno: %v, expected: %+v, actual: %+v", i, testcase.expected, actual)
	// 	}
	// }
}
