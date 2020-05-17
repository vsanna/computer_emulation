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

func (register *Register) String() string {
	return register.word.String()
}

func (register *Register) ToInt() int {
	n := 0
	for idx, b := range register.word.GetPrevious().Bits {
		n = n + b.GetVal()
		if idx != len(register.word.GetPrevious().Bits)-1 {
			n = n << 1
		}
	}

	return n
}
