package main

import (
	"computer_emulation/src/computer"
	"fmt"
)

func main() {
	fmt.Println("start")
	machine := computer.NewComputer()
	machine.Run()
}
