package vm_parser

import (
	"computer_emulation/src/vm/vm_ast"
	"computer_emulation/src/vm/vm_tokenizer"
	"strings"
	"testing"
)

func TestParser_ParseProgram_LineNumber(t *testing.T) {
	input := `
push constant 1
push constant 2
add
`
	t2 := vm_tokenizer.New(input)
	parser := New(t2)
	program := parser.ParseProgram()
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't have 25 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedLineNumber int
	}{
		{expectedLineNumber: 0},
		{expectedLineNumber: 1},
		{expectedLineNumber: 2},
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

func TestParser_ParseProgram_AddStatement(t *testing.T) {
	tokenizer := vm_tokenizer.New("add")
	parser := New(tokenizer)
	program := parser.ParseProgram()

	if len(parser.Errors()) > 0 {
		t.Fatal("program.Statements has errors.\n" + strings.Join(parser.Errors(), "\n"))
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements has more than 1 statement. got=%d", len(program.Statements))
	}

	statement := program.Statements[0]

	if statement.String() != "ADD" {
		t.Fatalf("statement.String() is wrong. expected = %q, actual = %q", "ADD", statement.String())
	}

	_, ok := statement.(*vm_ast.AddStatement)

	if !ok {
		t.Fatalf("failed to convert from Statement to AllocationStatement. actual=%T", statement)
	}
}
