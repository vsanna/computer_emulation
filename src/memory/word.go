package memory

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"fmt"
	"strconv"
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
	if in == nil {
		in = NewBus(BusOption{})
	}

	out := NewBus(BusOption{})
	for idx := range word.registers {
		bit := in.Bits[idx]
		out.Bits[idx] = word.registers[idx].Pass(bit, load)
	}
	return out
}

// NOTE: 実際には一つ前のtickでinした値を再現するための関数
func (word *Word) GetPrevious() *Bus {
	return word.Pass(nil, OFF)
}

func (word *Word) String() string {
	str := ""
	for idx, _ := range word.registers {
		str += word.registers[idx].String()
	}
	return str
}

func (word *Word) Load(instruction string) {
	if len(instruction) != WORD_WIDTH {
		panic(fmt.Sprintf("instruction is too long. max bit length per one instruction is %d", WORD_WIDTH))
	}

	for idx, _ := range instruction {
		rawbit := string([]rune(instruction)[idx])
		b, err := strconv.ParseInt(rawbit, 10, 4)
		if err != nil {
			panic("something wrong")
		}

		word.registers[idx].dff.Pass(bit.ToBit((int)(b)))
	}
}
