package vm_ast

import (
	"computer_emulation/src/vm/vm_tokenizer"
)

type Statement interface {
	statement()
	String() string
	LineNumber() int
}

type PushStatement struct {
	Segment vm_tokenizer.Token
	Value   vm_tokenizer.Token
	Line    int
}

func NewPushStatement() *PushStatement {
	return &PushStatement{
		Value:   vm_tokenizer.ZeroToken(),
		Segment: vm_tokenizer.ZeroToken(),
	}
}

func (a *PushStatement) statement() {
}

func (a *PushStatement) String() string {
	result := vm_tokenizer.PUSH
	if !a.Segment.IsZero() {
		result += " " + a.Segment.Literal
	}
	result += " " + a.Value.Literal
	return result
}

func (a *PushStatement) LineNumber() int {
	return a.Line
}

type AddStatement struct {
	Line int
}

func (a AddStatement) statement() {
}

func (a AddStatement) String() string {
	return vm_tokenizer.ADD
}

func (a AddStatement) LineNumber() int {
	return a.Line
}
