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
	memory                *memory.Memory
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

func NewCpu(mem *memory.Memory) *Cpu {
	return &Cpu{
		memory:                mem,
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
	log.Printf("[DEBUG] isCommandA=%v, address=%v, opsBus=%v, destBus=%v, jumpBus=%v\n",
		isCommandA, address, opsBus, destBus, jumpBus,
	)
	// 2-1. commandA: update A register
	cpu.a_reg.Pass(address, isCommandA)
	// 2-2-1. commandB: prepare commandB
	xBus, yBus, opsCodeBus := cpu.comp.Pass(
		cpu.a_reg.Pass(nil, OFF),
		cpu.d_reg.Pass(nil, OFF),
		cpu.memory.Pass(nil, OFF, cpu.a_reg.Pass(nil, OFF)),
		opsBus,
	)
	// 2-2-2. commandB: calculation
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
	cpu.memory.Pass(aluOutput, loadM, cpu.a_reg.Pass(nil, OFF))
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

func (cpu *Cpu) ShowDebugInfo() {
	log.Printf(
		"[DEBUG] A=%v, D=%v, M=%v, PC=%v, INST=%v\n",
		cpu.a_reg.ToInt(),
		cpu.d_reg.ToInt(),
		cpu.memory.Pass(nil, OFF, cpu.a_reg.Pass(nil, OFF)),
		cpu.pc_reg.ToInt(),
		cpu.memory.Pass(nil, OFF, cpu.pc_reg.Pass(nil, OFF)),
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
