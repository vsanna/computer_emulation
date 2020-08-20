package memory

import (
	. "computer_emulation/src/bit"
	"fmt"
	"strings"
)

/*
# Memory Mapping
This hardware architecture uses 16bit address-bus, so the max size of address is 2**16 = 65536
And it has two memory-space, and both starts its address with 0.
map:
- 0 - (2^15-1) : program memory(RO)
- 0 - (2^15-1) : data memory
	- 0 - 4 : predefined virtual registers for specific usage
	- 0 - 15: predefined virtual registers for general usage
    - 16 ~  : memory space for symbol tables(TODO: add validation system to restrict its reachable address)

TODO: consider how to use 1memory brock for two memory space
*/

const MEMORY_SIZE = 1024 * 64

//// predefined memory slots
//const SP_WORD_ADDRESS = 0
//const LCL_WORD_ADDRESS = 1
//const ARG_WORD_ADDRESS = 2
//const THIS_WORD_ADDRESS = 3
//const THAT_WORD_ADDRESS = 4
//const R0_WORD_ADDRESS = 0
//const R1_WORD_ADDRESS = 1
//const R2_WORD_ADDRESS = 2
//const R3_WORD_ADDRESS = 3
//const R4_WORD_ADDRESS = 4
//const R5_WORD_ADDRESS = 5
//const R6_WORD_ADDRESS = 6
//const R7_WORD_ADDRESS = 7
//const R8_WORD_ADDRESS = 8
//const R9_WORD_ADDRESS = 9
//const R10_WORD_ADDRESS = 10
//const R11_WORD_ADDRESS = 11
//const R12_WORD_ADDRESS = 12
//const R13_WORD_ADDRESS = 13
//const R14_WORD_ADDRESS = 14
//const R15_WORD_ADDRESS = 15
//
const SYMBOL_ENV_BASE_ADDRESS = 16 //16

type Memory struct {
	words []*Word
}

func NewMemory() *Memory {
	words := make([]*Word, MEMORY_SIZE+1)
	for i := 0; i <= MEMORY_SIZE; i++ {
		words[i] = NewWord()
	}
	return &Memory{words: words}
}

// load:1で新しい値に上書き(writeモード)
func (memory *Memory) Pass(in *Bus, load *Bit, address *Bus) *Bus {
	// TODO: cheating here by using ToInt.. should replace this logic
	addressInt := address.ToInt()
	return memory.words[addressInt].Pass(in, load)
}

// NOTE: This simulates ROM by loading machine language program from text file.
func (memory *Memory) LoadExecutable(machine_lang_program string) {
	lines := strings.Split(machine_lang_program, "\n")

	maxProgramLineNum := (MEMORY_SIZE)
	if len(lines) > maxProgramLineNum {
		panic(fmt.Sprintf("program is too long. max line number is %d", maxProgramLineNum))
	}

	for idx, line := range lines {
		memory.words[idx].Load(line)
	}
}
