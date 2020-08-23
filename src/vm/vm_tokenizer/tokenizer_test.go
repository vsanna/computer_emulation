package vm_tokenizer

import (
	"testing"
)

func TestTokenizer_NextToken(t *testing.T) {
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

label testline
goto testline
if_goto testline

function hoge 1
return
call hoge 2
`
	tokenizer := New(input)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{PUSH, "push"}, {CONSTANT, "constant"}, {INT, "2"},
		{PUSH, "push"}, {CONSTANT, "constant"}, {INT, "3"},
		{ADD, "add"},

		{PUSH, "push"}, {STATIC, "static"}, {INT, "0"},
		{PUSH, "push"}, {LOCAL, "local"}, {INT, "0"},
		{PUSH, "push"}, {ARGUMENT, "argument"}, {INT, "0"},
		{PUSH, "push"}, {THIS, "this"}, {INT, "0"},
		{PUSH, "push"}, {THAT, "that"}, {INT, "0"},
		{PUSH, "push"}, {POINTER, "pointer"}, {INT, "0"},
		{PUSH, "push"}, {TEMP, "temp"}, {INT, "0"},
		{PUSH, "push"}, {CONSTANT, "constant"}, {INT, "0"},

		{SUB, "sub"},
		{NOT, "not"},
		{AND, "and"},
		{OR, "or"},
		{EQ, "eq"},
		{NEG, "neg"},
		{GT, "gt"},
		{LT, "lt"},

		{POP, "pop"}, {STATIC, "static"}, {INT, "0"},
		{POP, "pop"}, {LOCAL, "local"}, {INT, "0"},
		{POP, "pop"}, {ARGUMENT, "argument"}, {INT, "0"},
		{POP, "pop"}, {THIS, "this"}, {INT, "0"},
		{POP, "pop"}, {THAT, "that"}, {INT, "0"},
		{POP, "pop"}, {POINTER, "pointer"}, {INT, "0"},
		{POP, "pop"}, {TEMP, "temp"}, {INT, "0"},
		{POP, "pop"}, {CONSTANT, "constant"}, {INT, "0"},

		{LABEL, "label"}, {IDENT, "testline"},
		{GOTO, "goto"}, {IDENT, "testline"},
		{IFGOTO, "if_goto"}, {IDENT, "testline"},

		{FUNCTION, "function"}, {IDENT, "hoge"}, {INT, "1"},
		{RETURN, "return"},
		{CALL, "call"}, {IDENT, "hoge"}, {INT, "2"},

		{EOF, ""},
	}

	for i, tt := range tests {
		tok := tokenizer.NextToken()

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral is wrong. expected: %q, actual: %q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType is wrong. expected: %q, actual: %q", i, tt.expectedType, tok.Type)
		}
	}
}
