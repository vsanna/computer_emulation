package memory

import (
	. "computer_emulation/src/bit"
	"fmt"
	"strings"
)

/*
# Memory Mapping
This hardware architecture uses 16bit address-bus, so the max size of address is 2**16 = 65536
map:
- 0 - (2^15-1)    : code memory(RO)
- 2^15 - (2^16-1) : data memory
	- DATA_MEMORY_BASE + 0 - 4: predefined virtual registers for specific usage
	- DATA_MEMORY_BASE + 0 - 15: predefined virtual registers for general usage
    - DATA_MEMORY_BASE + 16 ~  : memory space for symbol tables(TODO: add validation system to restrict its reachable address)
*/

const MEMORY_SIZE = 1024 * 64

const PROGRAM_MEMORY_BASE = 0
const DATA_MEMORY_BASE = 1024*32 - 1

// predefined memory slots
const SP_WORD_ADDRESS = DATA_MEMORY_BASE + 0
const LCL_WORD_ADDRESS = DATA_MEMORY_BASE + 1
const ARG_WORD_ADDRESS = DATA_MEMORY_BASE + 2
const THIS_WORD_ADDRESS = DATA_MEMORY_BASE + 3
const THAT_WORD_ADDRESS = DATA_MEMORY_BASE + 4
const R0_WORD_ADDRESS = DATA_MEMORY_BASE + 0
const R1_WORD_ADDRESS = DATA_MEMORY_BASE + 1
const R2_WORD_ADDRESS = DATA_MEMORY_BASE + 2
const R3_WORD_ADDRESS = DATA_MEMORY_BASE + 3
const R4_WORD_ADDRESS = DATA_MEMORY_BASE + 4
const R5_WORD_ADDRESS = DATA_MEMORY_BASE + 5
const R6_WORD_ADDRESS = DATA_MEMORY_BASE + 6
const R7_WORD_ADDRESS = DATA_MEMORY_BASE + 7
const R8_WORD_ADDRESS = DATA_MEMORY_BASE + 8
const R9_WORD_ADDRESS = DATA_MEMORY_BASE + 9
const R10_WORD_ADDRESS = DATA_MEMORY_BASE + 10
const R11_WORD_ADDRESS = DATA_MEMORY_BASE + 11
const R12_WORD_ADDRESS = DATA_MEMORY_BASE + 12
const R13_WORD_ADDRESS = DATA_MEMORY_BASE + 13
const R14_WORD_ADDRESS = DATA_MEMORY_BASE + 14
const R15_WORD_ADDRESS = DATA_MEMORY_BASE + 15

const SYMBOL_ENV_BASE_ADDRESS = 1024 //DATA_MEMORY_BASE + 16

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

	maxProgramLineNum := (MEMORY_SIZE - DATA_MEMORY_BASE)
	if len(lines) > maxProgramLineNum {
		panic(fmt.Sprintf("program is too long. max line number is %d", maxProgramLineNum))
	}

	for idx, line := range lines {
		memory.words[PROGRAM_MEMORY_BASE+idx].Load(line)
	}
}
