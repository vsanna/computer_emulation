package computer

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/cpu"
	"computer_emulation/src/memory"
	"log"
	"os"
	"time"
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
	log.Printf("computer starts running....\n")
	// 1. load ROM from txt file
	program := "0000111111111111\n" + // @4095 / SET A 4095
		"0000000000001101\n" + // @13 / SET A 13
		"1110111111010000" // D=1 / SET D 1
	computer.memory.Load(program)

	// 2. run
	// use infinite for loop instead of clock
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
		time.Sleep(200 * time.Millisecond)
	}
}

func (computer *Computer) ticktack(reset *Bit) {
	computer.cpu.ShowDebugInfo()
	// TODO: 図をここに貼る
	cpuOutput, resetBit := computer.cpu.StartTicktack(reset)
	memoryOutput := computer.memory.Pass(nil, OFF, cpuOutput)
	computer.cpu.Pass(memoryOutput, resetBit)
}
