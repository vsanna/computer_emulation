package assembler

import (
	"computer_emulation/src/assembler/parser"
	"computer_emulation/src/assembler/tokenizer"
	"computer_emulation/src/assembler/translator"
	"github.com/labstack/gommon/log"
	"io/ioutil"
)

// from file or string
type Assembler struct {
}

func New() *Assembler {
	return &Assembler{}
}

func (a Assembler) FromString(assemblerCode string) string {
	t := tokenizer.New(assemblerCode)
	p := parser.New(t)
	program := p.ParseProgram()
	t2 := translator.New(program)
	return t2.Translate()
}

func (a Assembler) FromFile(filepath string) string {
	filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("%q", err)
	}
	return a.FromString(string(filecontent))
}
