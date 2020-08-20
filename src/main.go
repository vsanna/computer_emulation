package main

import (
	"computer_emulation/src/assembler"
	"computer_emulation/src/computer"
	"fmt"
	"log"
	"os"
)

func main() {
	runMachine()
	//asmToMachineCode()
}

func runMachine() {
	fmt.Println("[HARDWARE] start main process.")
	machine := computer.NewComputer()
	machine.Run()
}

func asmToMachineCode() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./sample_asm/countUpTo100.asm"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given and couldn't find ./countUpTo100.binary. please place or pass your .choco file")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	assm := assembler.New()
	program := assm.FromFile(filename)
	log.Print(program)
}
