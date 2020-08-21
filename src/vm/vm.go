package vm

import (
	"computer_emulation/src/vm/vm_parser"
	"computer_emulation/src/vm/vm_tokenizer"
	"computer_emulation/src/vm/vm_translater"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

// from file or string
type VM struct {
}

func New() *VM {
	return &VM{}
}

func (a VM) FromString(vmCode string) string {
	t := vm_tokenizer.New(vmCode)
	p := vm_parser.New(t)
	program := p.ParseProgram()
	t2 := vm_translater.New(program)
	return t2.Translate()
}

func (a VM) FromFile(filepath string) string {
	filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	code := string(filecontent)

	return a.FromString(code)
}
