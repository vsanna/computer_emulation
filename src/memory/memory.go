package memory

import (
	. "computer_emulation/src/bit"
	"fmt"
	"strings"
)

// 16bitバスなのでアドレスの個数上限が2**16 = 65536
// 2Byte/word * (2**16) = 128KB
// NOTE: 2**16 == 2**10 * 2**6
const MEMORY_LENGTH = 1024 * 64

// うち、機械語を載せられるtextareaは0 - 1023まで
const TEXTAREA_MAX_LINENUM = 1023

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
func (memory *Memory) LoadExecutable(machine_lang_program string) {
	lines := strings.Split(machine_lang_program, "\n")

	if len(lines) > TEXTAREA_MAX_LINENUM {
		panic(fmt.Sprintf("program is too long. max line length is %d", TEXTAREA_MAX_LINENUM))
	}

	for idx, line := range lines {
		memory.words[idx].Load(line)
	}
}
