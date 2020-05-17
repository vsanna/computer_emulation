package memory

import (
	. "computer_emulation/src/bit"
	"fmt"
	"strings"
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
	// TODO: このToIntはずる。
	addressInt := address.ToInt()
	return memory.words[addressInt].Pass(in, load)
}

// NOTE: This simulate ROM by loading machine language program from text file.
func (memory *Memory) Load(machine_lang_program string) {
	lines := strings.Split(machine_lang_program, "\n")

	if len(lines) > MEMORY_LENGTH {
		panic(fmt.Sprintf("program is too long. max line length is %d", MEMORY_LENGTH))
	}

	for idx, line := range lines {
		memory.words[idx].Load(line)
	}
}
