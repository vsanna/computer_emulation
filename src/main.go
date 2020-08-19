package main

import (
	"computer_emulation/src/computer"
	"fmt"
)

func main() {
	fmt.Println("[HARDWARE] start main process.")
	machine := computer.NewComputer()
	machine.Run()
}
