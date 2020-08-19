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

// ≒ mother board(circuit basis)
type Computer struct {
	memory *memory.Memory
	cpu    *cpu.Cpu
	assm   *assembler.Assembler
}

func NewComputer() *Computer {
	memory := memory.NewMemory()
	computer := &Computer{
		memory: memory,
		cpu:    cpu.NewCpu(memory),
		assm:   assembler.New(),
	}

	return computer
}

func (computer *Computer) Run() {
	log.Printf("[HARDWARE] computer starts running....\n")
	// 1. load ROM from txt file
	// this is corresponding to Booting phase of CPU.
	program := "0000111111111111\n" + // @4095 / SET A 4095
		"1110111111001000\n" + // M=1
		"1111110111001000\n" + // M=M+1
		"1111110111001000\n" + // M=M+1
		"0000000000001101\n" + // @13 / SET A 13
		"1110111111000000\n" + // どこにもセットしない
		"1110111111100000\n" + // A=1
		"1110111111010000\n" + // D=1
		"1110111111001000\n" + // M=1
		"1110111111110000\n" + // AD=1
		"1110111111011000\n" + // DM=1
		"1110111111101000\n" + // AM=1
		"1110111111111000" // ADM=1
	computer.memory.LoadExecutable(program)

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
		time.Sleep(200 * time.Millisecond)
	}
}

// TODO: 図をここに貼る
// TODO: debugコードをflagで出し分け
func (computer *Computer) ticktack(reset *Bit) {
	computer.cpu.ShowDebugInfo()
	cpuOutput, resetBit := computer.cpu.StartTicktack(reset)
	memoryOutput := computer.memory.Pass(nil, OFF, cpuOutput)
	computer.cpu.Pass(memoryOutput, resetBit)
}
