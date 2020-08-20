package computer

import (
	"computer_emulation/src/assembler"
	. "computer_emulation/src/bit"
	"computer_emulation/src/cpu"
	"computer_emulation/src/memory"
	"log"
	"os"
	"strings"
	"time"
)

// ≒ mother board(circuit basis)
type Computer struct {
	data_memory    *memory.Memory
	program_memory *memory.Memory // ROM
	cpu            *cpu.Cpu
	assm           *assembler.Assembler
}

func NewComputer() *Computer {
	data_memory := memory.NewMemory()
	program_memory := memory.NewMemory()
	computer := &Computer{
		data_memory:    data_memory,
		program_memory: program_memory,
		cpu:            cpu.NewCpu(data_memory, program_memory),
		assm:           assembler.New(),
	}

	return computer
}

func (computer *Computer) Run() {
	log.Printf("[HARDWARE] computer starts running....\n")
	// 1. place binary code in memory.
	// the first line(memory[0]) is a booting process.
	program := strings.TrimSpace(`
0000000000010000
1110111111001000
0000000000010001
1110101010001000
0000000000010000
1111110000010000
0000000000000101
1110010011010000
0000000000010010
1110001100000001
0000000000010000
1111110000010000
0000000000010001
1111000010001000
0000000000010000
1111110111001000
0000000000000100
1110101010000111
0000000000010010
1110101010000111
`)
	computer.program_memory.LoadExecutable(program)

	// 2. run
	// use infinite loop instead of clock
	for {
		// 1. update user input
		// simulate RESET action by checking whether reset.txt exists or not
		// TODO: これをbitで扱うべきか
		resetInput := OFF
		_, err := os.Stat("reset.txt")
		if err == nil {
			resetInput = ON
		}

		// 2. operate current ticktack
		computer.ticktack(resetInput)

		// 3. update user output

		// 4. for debugging:
		time.Sleep(100 * time.Millisecond)
	}
}

// TODO: 図をここに貼る
// TODO: debugコードをflagで出し分け
func (computer *Computer) ticktack(reset *Bit) {
	computer.cpu.ShowDebugInfoForStatus()
	pcAddress, resetBit := computer.cpu.StartTicktack(reset)
	inst := computer.program_memory.Pass(nil, OFF, pcAddress)
	computer.cpu.Pass(inst, resetBit)
}
