package memory

import (
	. "computer_emulation/src/bit"
)

// 16bitバスなのでアドレスの個数上限が65536
// 2B * 1024 * 64 = 128KB
const MEMORY_LENGTH = 1024 * 64

type Memory struct {
	words []*Word
}

func NewMemory() *Memory {
	words := make([]*Word, MEMORY_LENGTH)
	for i := 0; i < MEMORY_LENGTH; i++ {
		words[i] = NewWord()
	}
	return &Memory{words: words}
}

// load:1で新しい値に上書き(writeモード)
func (memory *Memory) Pass(in *Bus, load *Bit, address *Bus) *Bus {
	addressInt := 1 // TODO: Bus -> intを行う回路
	return memory.words[addressInt].Pass(in, load)
}
