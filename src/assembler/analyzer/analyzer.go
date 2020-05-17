package analyzer

import (
	"fmt"
	"interpreter/src/runner"
	"os"
)

func main() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./main.asm"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given and couldn't find ./main.asm. please place or pass your .choco file")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	runner.Run(filename, os.Stdin, os.Stdout)
}
