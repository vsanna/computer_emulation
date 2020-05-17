package register

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/memory"
)

type Register struct {
	word *memory.Word
	name string
}

func NewRegister(name string) *Register {
	return &Register{word: memory.NewWord(), name: name}
}

func (register *Register) Pass(in *Bus, load *Bit) *Bus {
	return register.word.Pass(in, load)
}
