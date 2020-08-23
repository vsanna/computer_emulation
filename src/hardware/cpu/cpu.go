package cpu

import (
	"computer_emulation/src/hardware/alu"
	. "computer_emulation/src/hardware/bit"
	"computer_emulation/src/hardware/gate"
	"computer_emulation/src/hardware/memory"
	"computer_emulation/src/hardware/register"
	"log"
)

type Cpu struct {
	data_memory           *memory.Memory
	program_memory        *memory.Memory // ROM
	alu                   *alu.Alu
	decoder               *Decoder
	comp                  *alu.Comp
	dest                  *alu.Dest
	jump                  *alu.Jump
	a_reg                 *register.Register
	d_reg                 *register.Register
	pc_reg                *register.Register
	reset_bit             *memory.RegisterCell
	multibit_multi_plexer *gate.MultibitMultiPlexer
}

func NewCpu(data_memo *memory.Memory, program_memory *memory.Memory) *Cpu {
	return &Cpu{
		data_memory:           data_memo,
		program_memory:        program_memory,
		alu:                   alu.NewAlu(),
		decoder:               NewDecoder(),
		comp:                  alu.NewComp(),
		dest:                  alu.NewDest(),
		jump:                  alu.NewJump(),
		a_reg:                 register.NewRegister("A"),
		d_reg:                 register.NewRegister("D"),
		pc_reg:                register.NewRegister("PC"),
		reset_bit:             memory.NewRegisterCell(),
		multibit_multi_plexer: gate.NewMultibitMultiPlexer(),
	}
}

func (cpu *Cpu) StartTicktack(reset *Bit) (*Bus, *Bit) {
	return cpu.pc_reg.Pass(nil, OFF), cpu.reset_bit.Pass(reset, ON)
}

// TODO: refactoring
// an instruction goes through CPU and that updates registers and memory.
func (cpu *Cpu) Pass(in *Bus, resetBit *Bit) {
	// 1. decode instruction
	// instruction -> (atValue | isCommandA, opsBus, destBus, jumpBus)
	isCommandA, atValue, opsBus, destBus, jumpBus := cpu.decode(in)
	// cpu.ShowDebugInfoForOperation(atValue, isCommandA, opsBus, destBus, jumpBus)

	// 2-1. (if isCommandA) update A register as A-command
	cpu.a_reg.Pass(atValue, isCommandA)

	// 2-2. (if !isCommandA) update registers/memories as B-command
	// 2-2-1. commandB: prepare commandB
	xBus, yBus, opsCodeBus := cpu.comp.Pass(
		cpu.a_reg.Pass(nil, OFF),
		cpu.d_reg.Pass(nil, OFF),
		cpu.data_memory.Pass(nil, OFF, cpu.a_reg.Pass(nil, OFF)),
		opsBus,
	)
	// 2-2-2. commandB: retrieve alu result
	// passing 2 16bit bus and 6 flags as input
	aluOutput, outputIsZero, outputIsNegative := cpu.alu.Pass(
		xBus,
		yBus,
		opsCodeBus.Bits[0],
		opsCodeBus.Bits[1],
		opsCodeBus.Bits[2],
		opsCodeBus.Bits[3],
		opsCodeBus.Bits[4],
		opsCodeBus.Bits[5],
	)

	// 2-2-3. commandB: update register/memory
	loadA, loadD, loadM := cpu.dest.Pass(destBus)
	cpu.data_memory.Pass(aluOutput, loadM, cpu.a_reg.Pass(nil, OFF))
	cpu.a_reg.Pass(aluOutput, loadA)
	cpu.d_reg.Pass(aluOutput, loadD)

	// 2-3. update pc registry
	cpu.pc_reg.Pass(
		cpu.jump.Pass(
			cpu.a_reg.Pass(nil, OFF),
			cpu.pc_reg.Pass(nil, OFF),
			outputIsZero,
			outputIsNegative,
			jumpBus,
			resetBit,
		),
		ON,
	)

	return
}

func (cpu *Cpu) Reset(resetBit *Bit) {
	cpu.pc_reg.Pass(NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}), resetBit)
	return
}

func (cpu *Cpu) decode(in *Bus) (
	isCommandA *Bit,
	address *Bus,
	opsBus *Bus,
	destBus *Bus,
	jumpBus *Bus,
) {
	return cpu.decoder.Pass(in)
}

/**********************************
* for debugging
**********************************/

func (cpu *Cpu) ShowDebugInfoForStatus() {
	log.Printf(
		"[DEBUG] A=%v, D=%v, M(Memory[A])=%v, PC=%v:NEXT_INST=%v\n",
		cpu.a_reg.ToInt(),
		cpu.d_reg.ToInt(),
		cpu.data_memory.Pass(nil, OFF, cpu.a_reg.Pass(nil, OFF)).ToInt(),
		cpu.pc_reg.ToInt(),
		cpu.program_memory.Pass(nil, OFF, cpu.pc_reg.Pass(nil, OFF)),
	)
	cpu.ShowDebugInfoForGlobalStack()
	cpu.ShowDebugInfoForSegments()
}

func (cpu *Cpu) ShowDebugInfoForGlobalStack() {
	globalStack := []int{}
	start := memory.GLOBAL_STACK_BASE_ADDRESS
	end := cpu.data_memory.Pass(nil, OFF, IntToBus(memory.SP_WORD_ADDRESS)).ToInt()
	for i := start; i <= end; i++ {
		globalStack = append(globalStack, cpu.data_memory.Pass(nil, OFF, IntToBus(i)).ToInt())
	}
	// NOTE: last pos is always 0 since it's where new val comes in.
	log.Printf("[DEBUG] GlobalStack = %v", globalStack)
}

// NOTE:
// this/that/argument/local: [base address, base address+5]
// static/temp/pointer
func (cpu *Cpu) ShowDebugInfoForSegments() {
	localSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.LCL_WORD_ADDRESS)).ToInt()+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.LCL_WORD_ADDRESS)).ToInt()+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.LCL_WORD_ADDRESS)).ToInt()+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.LCL_WORD_ADDRESS)).ToInt()+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.LCL_WORD_ADDRESS)).ToInt()+4)).ToInt(),
	}
	argumentSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.ARG_WORD_ADDRESS)).ToInt()+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.ARG_WORD_ADDRESS)).ToInt()+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.ARG_WORD_ADDRESS)).ToInt()+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.ARG_WORD_ADDRESS)).ToInt()+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.ARG_WORD_ADDRESS)).ToInt()+4)).ToInt(),
	}
	pointerSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt(),
	}
	thisSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt()+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt()+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt()+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt()+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THIS_WORD_ADDRESS)).ToInt()+4)).ToInt(),
	}
	thatSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt()+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt()+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt()+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt()+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(cpu.data_memory.Pass(nil, OFF, IntToBus(memory.THAT_WORD_ADDRESS)).ToInt()+4)).ToInt(),
	}
	staticSegument := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.STATIC_BASE_ADDRESS+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.STATIC_BASE_ADDRESS+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.STATIC_BASE_ADDRESS+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.STATIC_BASE_ADDRESS+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.STATIC_BASE_ADDRESS+4)).ToInt(),
	}
	tempSegment := []int{
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.TEMP0_WORD_ADDRESS+0)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.TEMP0_WORD_ADDRESS+1)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.TEMP0_WORD_ADDRESS+2)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.TEMP0_WORD_ADDRESS+3)).ToInt(),
		cpu.data_memory.Pass(nil, OFF, IntToBus(memory.TEMP0_WORD_ADDRESS+4)).ToInt(),
	}

	log.Printf("[DEBUG] LOCAL    = %v", localSegment)
	log.Printf("[DEBUG] ARGUMENT = %v", argumentSegment)
	log.Printf("[DEBUG] POINTER  = %v", pointerSegment)
	log.Printf("[DEBUG] THIS     = %v", thisSegment)
	log.Printf("[DEBUG] THAT     = %v", thatSegment)
	log.Printf("[DEBUG] STATIC   = %v", staticSegument)
	log.Printf("[DEBUG] TEMP     = %v", tempSegment)

	dataMemoryHeadsSegment := []int{}
	for i := 0; i < 10; i++ {
		dataMemoryHeadsSegment = append(dataMemoryHeadsSegment, cpu.data_memory.Pass(nil, OFF, IntToBus(i)).ToInt())
	}
	log.Printf("[DEBUG] MEMORY   = %v", dataMemoryHeadsSegment)
}

func (cpu *Cpu) ShowDebugInfoForOperation(address *Bus, isCommandA *Bit, opsBus *Bus, destBus *Bus, jumpBus *Bus) {
	log.Printf("[DEBUG] inst=%v => isCommandA=%v opsBus=%v, destBus=%v, jumpBus=%v\n",
		address,
		isCommandA,
		[]*Bit{opsBus.Bits[0], opsBus.Bits[1], opsBus.Bits[2], opsBus.Bits[3], opsBus.Bits[4], opsBus.Bits[5], opsBus.Bits[6]},
		[]*Bit{destBus.Bits[0], destBus.Bits[1], destBus.Bits[2]},
		[]*Bit{jumpBus.Bits[0], jumpBus.Bits[1], jumpBus.Bits[2]},
	)
}
