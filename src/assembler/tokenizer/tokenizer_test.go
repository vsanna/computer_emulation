package tokenizer

import (
	"testing"
)

func TestTokenizer_NextToken(t *testing.T) {
	input := `
	@i
	M=1;
	@sum
	M=0;
(LOOP)
	@i
	D=M;
	@100
	D=D-A;
	@END
	D;JGT
	@i
	D=M;
	@sum
	M=D+M;
	@i
	M=M+1;
	@LOOP
	0;JMP
(END)
	@END
	0;JMP
	@DMAN
	@ADMA
`
	tokenizer := New(input)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{AT, "@"},
		{IDENT, "i"},
		{MEMORY, "M"},
		{ASSIGN, "="},
		{INT, "1"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{IDENT, "sum"},
		{MEMORY, "M"},
		{ASSIGN, "="},
		{INT, "0"},
		{SEMICOLON, ";"},
		{LPAREN, "("},
		{IDENT, "LOOP"},
		{RPAREN, ")"},
		{AT, "@"},
		{IDENT, "i"},
		{D_REG, "D"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{INT, "100"},
		{D_REG, "D"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{MINUS, "-"},
		{A_REG, "A"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{IDENT, "END"},
		{D_REG, "D"},
		{SEMICOLON, ";"},
		{JGT, "JGT"},
		{AT, "@"},
		{IDENT, "i"},
		{D_REG, "D"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{IDENT, "sum"},
		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{PLUS, "+"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{IDENT, "i"},
		{MEMORY, "M"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{PLUS, "+"},
		{INT, "1"},
		{SEMICOLON, ";"},
		{AT, "@"},
		{IDENT, "LOOP"},
		{INT, "0"},
		{SEMICOLON, ";"},
		{JMP, "JMP"},
		{LPAREN, "("},
		{IDENT, "END"},
		{RPAREN, ")"},
		{AT, "@"},
		{IDENT, "END"},
		{INT, "0"},
		{SEMICOLON, ";"},
		{JMP, "JMP"},

		{AT, "@"},
		{IDENT, "DMAN"},

		{AT, "@"},
		{IDENT, "ADMA"},

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

func TestTokenizer_NextToken2(t *testing.T) {
	input := `
	M=0;
	M=1;
	M=A;
	M=D;
	M=M;

	M=-1;
	M=-A;
	M=-D;
	M=-M;
	M=!A;
	M=!D;
	M=!M;

	M=A+1;
	M=D+1;
	M=M+1;

	M=A-1;
	M=D-1;
	M=M-1;

	M=A+D;
	M=D+M;

	M=A-D;
	M=D-A;
	M=D-M;
	M=M-D;
	
	M=A&D;
	M=D&M;
	
	M=A|D;
	M=D|M;
`
	//AとMの演算子はない。

	tokenizer := New(input)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{MEMORY, "M"},
		{ASSIGN, "="},
		{INT, "0"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MINUS, "-"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MINUS, "-"},
		{A_REG, "A"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MINUS, "-"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MINUS, "-"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{BANG, "!"},
		{A_REG, "A"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{BANG, "!"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{BANG, "!"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{PLUS, "+"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{PLUS, "+"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{PLUS, "+"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{MINUS, "-"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{MINUS, "-"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{MINUS, "-"},
		{INT, "1"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{PLUS, "+"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{PLUS, "+"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{MINUS, "-"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{MINUS, "-"},
		{A_REG, "A"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{MINUS, "-"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{MEMORY, "M"},
		{MINUS, "-"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{AND, "&"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{AND, "&"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{A_REG, "A"},
		{OR, "|"},
		{D_REG, "D"},
		{SEMICOLON, ";"},

		{MEMORY, "M"},
		{ASSIGN, "="},
		{D_REG, "D"},
		{OR, "|"},
		{MEMORY, "M"},
		{SEMICOLON, ";"},

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
