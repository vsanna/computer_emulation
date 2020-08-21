package vm_parser

import (
	"computer_emulation/src/vm/vm_ast"
	"computer_emulation/src/vm/vm_tokenizer"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

type Parser struct {
	tokenizer    *vm_tokenizer.Tokenizer
	currentToken vm_tokenizer.Token
	peekToken    vm_tokenizer.Token
	errors       []string
	currentLine  int
}

func New(tokenizer *vm_tokenizer.Tokenizer) *Parser {
	p := &Parser{tokenizer: tokenizer}
	p.errors = []string{}
	// 初期化
	p.nextToken() // peekTokenに一個目
	p.nextToken() // peekTokenに二個目, currentTokenに一個目
	return p
}

func (p *Parser) ParseProgram() *vm_ast.Program {
	program := &vm_ast.Program{}
	program.Statements = []vm_ast.Statement{}

	for p.currentToken.Type != vm_tokenizer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(tt vm_tokenizer.TokenType) {
	msg := fmt.Sprintf("current token is: %q, expected next token is: %q, got %q(%q)", p.currentToken.Literal, tt, p.peekToken.Type, p.peekToken.Literal)
	p.addError(msg)
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.tokenizer.NextToken()
}

func (p *Parser) parseStatement() vm_ast.Statement {
	switch p.currentToken.Type {
	case vm_tokenizer.PUSH:
		return p.parsePushStatement()
	//case vm_tokenizer.POP:
	//	return p.parsePopStatement()
	case vm_tokenizer.ADD, vm_tokenizer.SUB,
		vm_tokenizer.AND, vm_tokenizer.OR, vm_tokenizer.NOT,
		vm_tokenizer.EQ, vm_tokenizer.NEG, vm_tokenizer.GT, vm_tokenizer.LT:
		return p.parseArithmeticLogicalStatement()
	//case vm_tokenizer.LABEL:
	//	return p.parseLabelStatement()
	//case vm_tokenizer.GOTO:
	//	return p.parseGotoStatement()
	//case vm_tokenizer.IFGOTO:
	//	return p.parseIfgotoStatement()
	//case vm_tokenizer.RETURN:
	//	return p.parseReturnStatement()
	//case vm_tokenizer.FUNCTION:
	//	return p.parseFunctionStatement()
	//case vm_tokenizer.CALL:
	//	return p.parseCallStatement()
	default:
		return p.parseAssignmentStatement()
	}
}

func (p *Parser) parsePushStatement() vm_ast.Statement {
	statement := vm_ast.NewPushStatement()

	if !(p.peekTokenIs(vm_tokenizer.IDENT) ||
		p.peekTokenIs(vm_tokenizer.INT) ||
		p.peekTokenIs(vm_tokenizer.CONSTANT) ||
		p.peekTokenIs(vm_tokenizer.ARGUMENT) ||
		p.peekTokenIs(vm_tokenizer.LOCAL) ||
		p.peekTokenIs(vm_tokenizer.STATIC) ||
		p.peekTokenIs(vm_tokenizer.THIS) ||
		p.peekTokenIs(vm_tokenizer.THAT) ||
		p.peekTokenIs(vm_tokenizer.POINTER) ||
		p.peekTokenIs(vm_tokenizer.TEMP) ||
		p.peekTokenIs(vm_tokenizer.POINTER)) {
		p.addError(fmt.Sprintf("unexpected token. actual=%q", p.currentToken.Type))
		log.Error("parse error")
		os.Exit(1)
		return nil
	}

	// be in [push] val
	if p.peekTokenIs(vm_tokenizer.IDENT) || p.peekTokenIs(vm_tokenizer.INT) {
		// move to push [val]
		p.nextToken()
		statement.Value = p.currentToken
		statement.Line = p.currentLine
		p.currentLine += 1
		p.nextToken()
		return statement
	}

	// be in [push] segment val
	// move to push [segment] val
	p.nextToken()
	statement.Segment = p.currentToken
	// move to push segment [val]
	p.nextToken()

	statement.Value = p.currentToken
	statement.Line = p.currentLine
	p.currentLine += 1
	p.nextToken()
	return statement
}

func (p *Parser) expectPeek(tt vm_tokenizer.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	} else {
		p.peekError(tt)
		return false
	}
}

func (p *Parser) peekTokenIs(tt vm_tokenizer.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) parseArithmeticLogicalStatement() vm_ast.Statement {
	// be in [add]
	var statement vm_ast.Statement
	switch p.currentToken.Type {
	case vm_tokenizer.ADD:
		statement = &vm_ast.AddStatement{Line: p.currentLine}
	case vm_tokenizer.SUB:
		statement = &vm_ast.SubStatement{Line: p.currentLine}
	case vm_tokenizer.AND:
		statement = &vm_ast.AndStatement{Line: p.currentLine}
	case vm_tokenizer.OR:
		statement = &vm_ast.OrStatement{Line: p.currentLine}

	case vm_tokenizer.EQ:
		statement = &vm_ast.EqStatement{Line: p.currentLine}
	case vm_tokenizer.GT:
		statement = &vm_ast.GtStatement{Line: p.currentLine}
	case vm_tokenizer.LT:
		statement = &vm_ast.LtStatement{Line: p.currentLine}

	case vm_tokenizer.NOT:
		statement = &vm_ast.NotStatement{Line: p.currentLine}
	case vm_tokenizer.NEG:
		statement = &vm_ast.NeqStatement{Line: p.currentLine}
	}

	p.currentLine += 1

	p.nextToken()

	return statement
}

// TODO: this is temp implementation for debug
func (p *Parser) parseAssignmentStatement() vm_ast.Statement {
	return &vm_ast.PushStatement{
		Value: vm_tokenizer.Token{
			Literal: "1",
			Type:    vm_tokenizer.INT,
		},
		Line: 1,
	}
}
