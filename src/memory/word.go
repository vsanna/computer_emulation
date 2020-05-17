package memory

import (
	. "computer_emulation/src/bit"
)

// 1word = 2Byte. 32bitマシンは4B
const WORD_WIDTH = 16

type Word struct {
	registers []*RegisterCell
}

func NewWord() *Word {
	registers := make([]*RegisterCell, WORD_WIDTH)
	for i := 0; i < WORD_WIDTH; i++ {
		registers[i] = NewRegisterCell()
	}
	return &Word{registers: registers}
}

func (word *Word) Pass(in *Bus, load *Bit) *Bus {
	out := NewBus(BusOption{})
	for idx, bit := range in.Bits {
		out.Bits[idx] = word.registers[idx].Pass(bit, load)
	}
	return out
}

// NOTE: 実際には一つ前のtickでinした値を再現するための関数
func (word *Word) GetPrevious() *Bus {
	return word.Pass(nil, OFF)
}
