package ast

import (
	"computer_emulation/src/assembler/tokenizer"
)

type Statement interface {
	statement()
	String() string
	LineNumber() int
}

type AllocationStatement struct {
	Value tokenizer.Token
	Line  int
}

func (a *AllocationStatement) statement() {
}

func (a *AllocationStatement) String() string {
	return tokenizer.AT + a.Value.Literal
}

func (a *AllocationStatement) LineNumber() int {
	return a.Line
}

type AddressTaggingStatement struct {
	Value tokenizer.Token
	Line  int
}

func (a *AddressTaggingStatement) statement() {
}

func (a *AddressTaggingStatement) LineNumber() int {
	return a.Line
}

func (a *AddressTaggingStatement) String() string {
	return tokenizer.LPAREN + a.Value.Literal + tokenizer.RPAREN
}

type OpsAndJumpStatement struct {
	Dest []tokenizer.Token
	Comp []tokenizer.Token
	Jump *tokenizer.Token
	Line int
}

func (o *OpsAndJumpStatement) statement() {
}

func (a *OpsAndJumpStatement) LineNumber() int {
	return a.Line
}

func (o *OpsAndJumpStatement) String() string {
	destString := ""
	for _, token := range o.Dest {
		destString += token.Literal
	}

	compString := ""
	for _, token := range o.Comp {
		compString += token.Literal
	}
	compString += tokenizer.SEMICOLON

	jumpString := ""
	if o.Jump != nil {
		jumpString += o.Jump.Literal
	}

	resultString := ""
	if len(destString) > 0 {
		resultString += destString + "="
	}
	resultString += compString
	resultString += jumpString

	return resultString
}
