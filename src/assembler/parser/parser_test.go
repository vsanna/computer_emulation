package parser

import (
	"computer_emulation/src/assembler/ast"
	"computer_emulation/src/assembler/tokenizer"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestParser_ParseProgram_LineNumber(t *testing.T) {
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
	M=M+1;
`
	t2 := tokenizer.New(input)
	parser := New(t2)
	program := parser.ParseProgram()
	if len(program.Statements) != 25 {
		t.Fatalf("program.Statements doesn't have 25 statements. got=%d", len(program.Statements))
	}

	for i, statement := range program.Statements {
		if statement.LineNumber() != i {
			t.Fatalf("unexpected LineNumber(). expected=%d, actual=%d", i, statement.LineNumber())
		}
	}
}

func TestParser_ParseProgram_AllocationStatement(t *testing.T) {
	tests := []struct {
		input                string
		expectedStatement    string
		expectedValueLiteral string
		expectedValueType    tokenizer.TokenType
	}{
		{"@i", "@i", "i", tokenizer.IDENT},
		{"@LOOP", "@LOOP", "LOOP", tokenizer.IDENT},
		{"@END", "@END", "END", tokenizer.IDENT},
		{"@100", "@100", "100", tokenizer.INT},
	}
	for _, test := range tests {
		tokenizer := tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements has more than 1 statement. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.expectedStatement {
			t.Fatalf("statement.String() is wrong. expected = %q, actual = %q", test.expectedStatement, statement.String())
		}

		allocStatement, ok := statement.(*ast.AllocationStatement)

		log.Print(allocStatement.String())

		if !ok {
			t.Fatalf("failed to convert from Statement to AllocationStatement. actual=%T", statement)
		}

		if allocStatement.Value.Literal != test.expectedValueLiteral {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueLiteral, allocStatement.Value.Literal)
		}
		if allocStatement.Value.Type != test.expectedValueType {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueType, allocStatement.Value.Type)
		}

	}
}

func TestParser_ParseProgram_AddressTaggingStatement(t *testing.T) {
	tests := []struct {
		input                string
		expectedStatement    string
		expectedValueLiteral string
		expectedValueType    tokenizer.TokenType
	}{
		{"(LOOP)", "(LOOP)", "LOOP", tokenizer.IDENT},
		{"(END)", "(END)", "END", tokenizer.IDENT},
		{"(MOONAND)", "(MOONAND)", "MOONAND", tokenizer.IDENT},
	}
	for _, test := range tests {
		tokenizer := tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()

		if len(parser.Errors()) > 0 {
			t.Fatal("program.Statements has errors.\n" + strings.Join(parser.Errors(), "\n"))
		}

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements has more than 1 statement. got=%d", len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.expectedStatement {
			t.Fatalf("statement.String() is wrong. expected = %q, actual = %q", test.expectedStatement, statement.String())
		}

		taggingStatement, ok := statement.(*ast.AddressTaggingStatement)

		log.Print(taggingStatement.String())

		if !ok {
			t.Fatalf("failed to convert from Statement to AllocationStatement. actual=%T", statement)
		}

		if taggingStatement.Value.Literal != test.expectedValueLiteral {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueLiteral, taggingStatement.Value.Literal)
		}
		if taggingStatement.Value.Type != test.expectedValueType {
			t.Fatalf("unexpected value. expected=%q, actual=%q", test.expectedValueType, taggingStatement.Value.Type)
		}

	}
}

func TestParser_ParseProgram_OpsAndJumpStatement(t *testing.T) {
	tests := []struct {
		input             string
		expectedDestTypes []tokenizer.TokenType
		expectedCompTypes []tokenizer.TokenType
		expectedJumpType  *tokenizer.Token
	}{
		{"M=M+1;",
			[]tokenizer.TokenType{tokenizer.MEMORY},
			[]tokenizer.TokenType{tokenizer.MEMORY, tokenizer.PLUS, tokenizer.INT},
			nil,
		},
		{"M=A-1;",
			[]tokenizer.TokenType{tokenizer.MEMORY},
			[]tokenizer.TokenType{tokenizer.A_REG, tokenizer.MINUS, tokenizer.INT},
			nil,
		},
		{"D;JGT",
			[]tokenizer.TokenType{},
			[]tokenizer.TokenType{tokenizer.D_REG},
			&tokenizer.Token{Type: tokenizer.JGT, Literal: "JGT"},
		},
		{"0;JMP",
			[]tokenizer.TokenType{},
			[]tokenizer.TokenType{tokenizer.INT},
			&tokenizer.Token{Type: tokenizer.JMP, Literal: "JMP"},
		},
		{"ADM=M-D;",
			[]tokenizer.TokenType{tokenizer.A_REG, tokenizer.D_REG, tokenizer.MEMORY},
			[]tokenizer.TokenType{tokenizer.MEMORY, tokenizer.MINUS, tokenizer.D_REG},
			nil,
		},
	}
	for i, test := range tests {
		tokenizer := tokenizer.New(test.input)
		parser := New(tokenizer)
		program := parser.ParseProgram()

		if len(parser.Errors()) > 0 {
			t.Fatal(strconv.Itoa(i) + "program.Statements has errors.\n" + strings.Join(parser.Errors(), "\n"))
		}

		if len(program.Statements) != 1 {
			t.Fatalf("[%d] program.Statements has more than 1 statement. got=%d", i, len(program.Statements))
		}

		statement := program.Statements[0]

		if statement.String() != test.input {
			t.Fatalf("[%d] statement.String() is wrong. expected = %q, actual = %q", i, test.input, statement.String())
		}

		opsAndJumpStatement, ok := statement.(*ast.OpsAndJumpStatement)

		if !ok {
			t.Fatalf("[%d] failed to convert from Statement to AllocationStatement. actual=%T", i, statement)
		}

		for ii, token := range opsAndJumpStatement.Dest {
			if token.Type != test.expectedDestTypes[ii] {
				t.Fatalf("[%d - %d] unexpected dest type. expected=%q, actual=%q", i, ii, test.expectedDestTypes[ii], token.Type)
			}
		}

		for ii, token := range opsAndJumpStatement.Comp {
			if token.Type != test.expectedCompTypes[ii] {
				t.Fatalf("[%d - %d] unexpected comp type. expected=%q, actual=%q", i, ii, test.expectedCompTypes[ii], token.Type)
			}
		}

		if test.expectedJumpType == nil {
			if opsAndJumpStatement.Jump != nil {
				t.Fatalf("[%d] unexpected jump token. expected=%q, actual=%q", i, test.expectedJumpType, opsAndJumpStatement.Jump)
			}
		} else {
			if opsAndJumpStatement.Jump == nil || opsAndJumpStatement.Jump.Type != test.expectedJumpType.Type {
				t.Fatalf("[%d] unexpected jump token. expected=%q, actual=%q", i, test.expectedJumpType, opsAndJumpStatement.Jump)
			}
		}

	}
}
