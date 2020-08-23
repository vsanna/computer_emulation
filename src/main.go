package main

import (
	"computer_emulation/src/assembler"
	"computer_emulation/src/hardware/computer"
	"computer_emulation/src/vm"
	"fmt"
	"log"
	"os"
)

func main() {
	//vmToAsm()
	//asmToMachineCode()
	//runMachine()

	runMachineFromVmCode() // for convenience
}

func vmToAsm() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./sample/vm/playground.vm"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	virtualMachine := vm.New()
	program := virtualMachine.FromFile(filename)
	log.Printf("\n%s", program)
}

func asmToMachineCode() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./sample/asm/sum_up_to_10.asm"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	assm := assembler.New()
	program := assm.FromFile(filename)
	log.Printf("\n%s", program)
}

func runMachine() {
	fmt.Println("[HARDWARE] start main process.")
	machine := computer.NewComputer()
	machine.Run()
}

// for convenience
// this translates vm file -> asm file -> binary file
func runMachineFromVmCode() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./sample/vm/func.vm"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	virtualMachine := vm.New()
	asmProgram := virtualMachine.FromFile(filename)

	assm := assembler.New()
	binaryProgram := assm.FromString(asmProgram)

	machine := computer.NewComputer()
	computer.LoadPresetBinaryProgram(machine)(binaryProgram)
	machine.Run()
}
