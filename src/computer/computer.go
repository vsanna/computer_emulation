package computer

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/cpu"
	"computer_emulation/src/memory"
	"os"
)

// ≒ mother board(circuit basis)
type Computer struct {
	memory *memory.Memory
	cpu    *cpu.Cpu
}

func NewComputer() *Computer {
	memory := memory.NewMemory()
	computer := &Computer{
		memory: memory,
		cpu:    cpu.NewCpu(memory),
	}

	return computer
}

func (computer *Computer) Run() {
	// 1. load ROM from txt file

	// 2. run
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
	}
}

func (computer *Computer) ticktack(reset *Bit) {
	// TODO: 図をここに貼る
	cpuOutput, resetBit := computer.cpu.StartTicktack(reset)
	memoryOutput := computer.memory.Pass(nil, OFF, cpuOutput)
	computer.cpu.Pass(memoryOutput, resetBit)
	return
}
