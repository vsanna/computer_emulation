package computer

import (
	"computer_emulation/src/assembler"
	. "computer_emulation/src/bit"
	"computer_emulation/src/cpu"
	"computer_emulation/src/memory"
	"log"
	"os"
	"time"
)

// ≒ mother board(circuit basis), which has cpu and memories on it.
type Computer struct {
	data_memory    *memory.Memory
	program_memory *memory.Memory // ROM
	cpu            *cpu.Cpu
}

func NewComputer() *Computer {
	data_memory := memory.NewMemory()
	program_memory := memory.NewMemory()
	computer := &Computer{
		data_memory:    data_memory,
		program_memory: program_memory,
		cpu:            cpu.NewCpu(data_memory, program_memory),
	}
	// place binary code in memory in advance
	// but for convenience, calling assembler here instead of pasting binary code
	program := assembler.New().FromFile("./sample_asm/func.asm")
	LoadPresetBinaryProgram(computer)(program)

	return computer
}

func LoadPresetBinaryProgram(computer *Computer) func(machine_lang_program string) {
	return computer.program_memory.LoadExecutable
}

func (computer *Computer) Run() {
	log.Printf("[HARDWARE] computer starts running....\n")

	// use infinite loop instead of clock
	for {
		// 1. update user input
		// simulate RESET action by checking whether reset.txt exists or not
		// TODO: consider whether to use resetInput as Bit type or not.
		resetInput := OFF
		_, err := os.Stat("reset.txt")
		if err == nil {
			resetInput = ON
		}

		// 2. operate current ticktack
		computer.ticktack(resetInput)

		// 3. update user output

		// 4. for debugging:
		time.Sleep(10 * time.Millisecond)
	}
}

// TODO: write and paste architecture diagram here
// TODO: control debug mode by flagging
func (computer *Computer) ticktack(reset *Bit) {
	computer.cpu.ShowDebugInfoForStatus()
	pcAddress, resetBit := computer.cpu.StartTicktack(reset)
	inst := computer.program_memory.Pass(nil, OFF, pcAddress)
	computer.cpu.Pass(inst, resetBit)
}
