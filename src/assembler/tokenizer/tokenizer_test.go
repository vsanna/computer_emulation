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
	@SP
	@LCL
	@ARG
	@THIS
	@THAT
	@R0
	@R1
	@R2
	@R3
	@R4
	@R5
	@R6
	@R7
	@R8
	@R9
	@R10
	@R11
	@R12
	@R13
	@R14
	@R15
	@generated_ident__0bd71bf4_e3ab_11ea_9852_acde48001122_THEN
	D;JEQ
	@generated_ident__0bd71bf4_e3ab_11ea_9852_acde48001122_ELSE
	0;JMP
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

		{AT, "@"}, {IDENT, "SP"},
		{AT, "@"}, {IDENT, "LCL"},
		{AT, "@"}, {IDENT, "ARG"},
		{AT, "@"}, {IDENT, "THIS"},
		{AT, "@"}, {IDENT, "THAT"},
		{AT, "@"}, {IDENT, "R0"},
		{AT, "@"}, {IDENT, "R1"},
		{AT, "@"}, {IDENT, "R2"},
		{AT, "@"}, {IDENT, "R3"},
		{AT, "@"}, {IDENT, "R4"},
		{AT, "@"}, {IDENT, "R5"},
		{AT, "@"}, {IDENT, "R6"},
		{AT, "@"}, {IDENT, "R7"},
		{AT, "@"}, {IDENT, "R8"},
		{AT, "@"}, {IDENT, "R9"},
		{AT, "@"}, {IDENT, "R10"},
		{AT, "@"}, {IDENT, "R11"},
		{AT, "@"}, {IDENT, "R12"},
		{AT, "@"}, {IDENT, "R13"},
		{AT, "@"}, {IDENT, "R14"},
		{AT, "@"}, {IDENT, "R15"},

		{AT, "@"}, {IDENT, "generated_ident__0bd71bf4_e3ab_11ea_9852_acde48001122_THEN"},
		{D_REG, "D"}, {SEMICOLON, ";"}, {JEQ, "JEQ"},

		{AT, "@"}, {IDENT, "generated_ident__0bd71bf4_e3ab_11ea_9852_acde48001122_ELSE"},
		{INT, "0"}, {SEMICOLON, ";"}, {JMP, "JMP"},

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
