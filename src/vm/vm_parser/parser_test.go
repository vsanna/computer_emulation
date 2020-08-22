package vm_parser

import (
	"computer_emulation/src/vm/vm_ast"
	"computer_emulation/src/vm/vm_tokenizer"
	"strings"
	"testing"
)

func TestParser_ParseProgram_LineNumber(t *testing.T) {
	input := `
push constant 2
push constant 3
add
push static 0
push local 0
push argument 0
push this 0
push that 0
push pointer 0
push temp 0
push constant 0
sub
not
and
or
eq
neg
gt
lt

pop static 0
pop local 0
pop argument 0
pop this 0
pop that 0
pop pointer 0
pop temp 0
pop constant 0
`
	t2 := vm_tokenizer.New(input)
	parser := New(t2)
	program := parser.ParseProgram()

	expectedLineNum := 27

	if len(program.Statements) != expectedLineNum {
		t.Fatalf("program.Statements doesn't have %d statements. got=%d", expectedLineNum, len(program.Statements))
	}

	tests := []struct {
		expectedLineNumber int
	}{
		{expectedLineNumber: 0},
		{expectedLineNumber: 1},
		{expectedLineNumber: 2},
		{expectedLineNumber: 3},
		{expectedLineNumber: 4},
		{expectedLineNumber: 5},
		{expectedLineNumber: 6},
		{expectedLineNumber: 7},
		{expectedLineNumber: 8},
		{expectedLineNumber: 9},
		{expectedLineNumber: 10},
		{expectedLineNumber: 11},
		{expectedLineNumber: 12},
		{expectedLineNumber: 13},
		{expectedLineNumber: 14},
		{expectedLineNumber: 15},
		{expectedLineNumber: 16},
		{expectedLineNumber: 17},
		{expectedLineNumber: 18},
		{expectedLineNumber: 19},
		{expectedLineNumber: 20},
		{expectedLineNumber: 21},
		{expectedLineNumber: 22},
		{expectedLineNumber: 23},
		{expectedLineNumber: 24},
		{expectedLineNumber: 25},
		{expectedLineNumber: 26},
	}

	for i, statement := range program.Statements {
		if statement.LineNumber() != tests[i].expectedLineNumber {
			t.Fatalf("unexpected LineNumber(). expected=%d, actual=%d, stmt = %s", tests[i].expectedLineNumber, statement.LineNumber(), statement)
		}
	}
}

func TestParser_ParseProgram_PushStatement(t *testing.T) {
	tests := []struct {
		input                string
		expectedStatement    string
		expectedValueLiteral string
		expectedValueType    vm_tokenizer.TokenType
	}{
		{"push constant 10", "PUSH constant 10", "10", vm_tokenizer.INT},
		{"push constant myval", "PUSH constant myval", "myval", vm_tokenizer.IDENT},
		{"push 10", "PUSH 10", "10", vm_tokenizer.INT},
		{"push myval", "PUSH myval", "myval", vm_tokenizer.IDENT},
	}
	for i, test := range tests {
		tokenizer := vm_tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements has more than 1 statement. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.expectedStatement {
			t.Fatalf("[%d] statement.String() is wrong. expected = %q, actual = %q", i, test.expectedStatement, statement.String())
		}

		pushStatement, ok := statement.(*vm_ast.PushStatement)

		if !ok {
			t.Fatalf("failed to convert from Statement to AllocationStatement. actual=%T", statement)
		}

		if pushStatement.Value.Literal != test.expectedValueLiteral {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueLiteral, pushStatement.Value.Literal)
		}
		if pushStatement.Value.Type != test.expectedValueType {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueType, pushStatement.Value.Type)
		}

	}
}

func TestParser_ParseProgram_PopStatement(t *testing.T) {
	tests := []struct {
		input                string
		expectedStatement    string
		expectedValueLiteral string
		expectedValueType    vm_tokenizer.TokenType
	}{
		{"pop constant 10", "POP constant 10", "10", vm_tokenizer.INT},
		{"pop constant myval", "POP constant myval", "myval", vm_tokenizer.IDENT},
		{"pop 10", "POP 10", "10", vm_tokenizer.INT},
		{"pop myval", "POP myval", "myval", vm_tokenizer.IDENT},
	}
	for i, test := range tests {
		tokenizer := vm_tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements has more than 1 statement. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.expectedStatement {
			t.Fatalf("[%d] statement.String() is wrong. expected = %q, actual = %q", i, test.expectedStatement, statement.String())
		}

		pushStatement, ok := statement.(*vm_ast.PopStatement)

		if !ok {
			t.Fatalf("failed to convert from Statement to AllocationStatement. actual=%T", statement)
		}

		if pushStatement.Value.Literal != test.expectedValueLiteral {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueLiteral, pushStatement.Value.Literal)
		}
		if pushStatement.Value.Type != test.expectedValueType {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueType, pushStatement.Value.Type)
		}

	}
}

func TestParser_parseArithmeticStatement(t *testing.T) {
	tests := []struct {
		input           string
		expectedLiteral string
	}{
		{"add", "ADD"},
		{"sub", "SUB"},
		{"and", "AND"},
		{"or", "OR"},

		{"eq", "EQ"},
		{"gt", "GT"},
		{"lt", "LT"},

		{"not", "NOT"},
		{"neg", "NEG"},
	}

	for i, test := range tests {
		tokenizer := vm_tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()

		if len(parser.Errors()) > 0 {
			t.Fatalf("[%d] program.Statements has errors.\n"+strings.Join(parser.Errors(), "\n"), i)
		}

		if len(program.Statements) != 1 {
			t.Fatalf("[%d] program.Statements has more than 1 statement. got=%d", i, len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.expectedLiteral {
			t.Fatalf("[%d] statement.String() is wrong. expected = %q, actual = %q", i, "ADD", statement.String())
		}
	}

}
