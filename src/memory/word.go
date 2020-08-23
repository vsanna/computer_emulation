package memory

import (
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"log"
	"strconv"
)

// This architecture uses 16bits-bus, so the word size is 2 bytes.
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

// for loading binary code in memory.
func (word *Word) Load(instruction string) {
	if len(instruction) != WORD_WIDTH {
		log.Fatalf(
			"instruction doesn't has proper length. bit length per one instruction should be %d, but got %d",
			WORD_WIDTH, len(instruction))
	}

	log.Printf("[DEBUG] loading program: %s\n", instruction)
	for idx, _ := range instruction {
		rawbit := string([]rune(instruction)[idx])
		b, err := strconv.ParseInt(rawbit, 10, 4)
		if err != nil {
			panic("something wrong")
		}

		word.registers[idx].dff.Pass(bit.ToBit((int)(b)))
	}
}

// For convenience to debug.
func (word *Word) GetPrevious() *Bus {
	return word.Pass(nil, OFF)
}

// For convenience to debug.
func (word *Word) String() string {
	str := ""
	for idx, _ := range word.registers {
		str += word.registers[idx].String()
	}
	return str
}
