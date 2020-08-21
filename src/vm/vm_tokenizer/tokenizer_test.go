package vm_tokenizer

import (
	"testing"
)

func TestTokenizer_NextToken(t *testing.T) {
	input := `
push constant 2
push constant 3
add
`
	tokenizer := New(input)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{PUSH, "push"},
		{CONSTANT, "constant"},
		{INT, "2"},
		{PUSH, "push"},
		{CONSTANT, "constant"},
		{INT, "3"},
		{ADD, "add"},
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
