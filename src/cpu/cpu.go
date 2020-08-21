package cpu

import (
	"computer_emulation/src/alu"
	"computer_emulation/src/bit"
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
	"computer_emulation/src/memory"
	"computer_emulation/src/register"
	"log"
)

/*register
address
data: 1 - 10
pc
*/

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
func (cpu *Cpu) Pass(in *Bus, resetBit *Bit) {
	// 1. decode
	isCommandA, address, opsBus, destBus, jumpBus := cpu.decode(in)
	//cpu.ShowDebugInfoForOperation(address, isCommandA, opsBus, destBus, jumpBus)

	// 2-1. commandA: update A register
	cpu.a_reg.Pass(address, isCommandA)

	// 2-2-1. commandB: prepare commandB
	// opsBusの内容に応じて、入力ソースをnull(0), A, D, Mの4つのうちから2つ選択する
	// ALUに応じた処理flagをopsCodeBusとして取得
	xBus, yBus, opsCodeBus := cpu.comp.Pass(
		cpu.a_reg.Pass(nil, OFF),
		cpu.d_reg.Pass(nil, OFF),
		cpu.data_memory.Pass(nil, OFF, cpu.a_reg.Pass(nil, OFF)),
		opsBus,
	)
	// 2-2-2. commandB: calculation
	// 2つの入力系統と、6つのflag
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

	// 2-3. update pc
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
	cpu.pc_reg.Pass(bit.NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}), resetBit)
	return
}

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
}

func (cpu *Cpu) ShowDebugInfoForGlobalStack() {
	globalStack := []int{}
	start := memory.GLOBAL_STACK_BASE_ADDRESS
	end := cpu.data_memory.Pass(nil, OFF, bit.IntToBus(memory.SP_WORD_ADDRESS)).ToInt()
	for i := start; i <= end; i++ {
		globalStack = append(globalStack, cpu.data_memory.Pass(nil, OFF, bit.IntToBus(i)).ToInt())
	}
	// NOTE: least pos is always 0 since it's where new val will come in.
	log.Printf("[DEBUG] GlobalStack = %v", start, end, globalStack)
}

func (cpu *Cpu) ShowDebugInfoForOperation(address *Bus, isCommandA *Bit, opsBus *Bus, destBus *Bus, jumpBus *Bus) {
	log.Printf("[DEBUG] inst=%v => isCommandA=%v opsBus=%v, destBus=%v, jumpBus=%v\n",
		address,
		isCommandA,
		[]*bit.Bit{opsBus.Bits[0], opsBus.Bits[1], opsBus.Bits[2], opsBus.Bits[3], opsBus.Bits[4], opsBus.Bits[5], opsBus.Bits[6]},
		[]*bit.Bit{destBus.Bits[0], destBus.Bits[1], destBus.Bits[2]},
		[]*bit.Bit{jumpBus.Bits[0], jumpBus.Bits[1], jumpBus.Bits[2]},
	)
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
