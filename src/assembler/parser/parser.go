package parser

import (
	"computer_emulation/src/assembler/ast"
	"computer_emulation/src/assembler/tokenizer"
	"fmt"
)

type Parser struct {
	tokenizer    *tokenizer.Tokenizer
	currentToken tokenizer.Token
	peekToken    tokenizer.Token
	errors       []string
	currentLine  int
}

func New(tokenizer *tokenizer.Tokenizer) *Parser {
	p := &Parser{tokenizer: tokenizer}
	p.errors = []string{}
	// 初期化
	p.nextToken() // peekTokenに一個目
	p.nextToken() // peekTokenに二個目, currentTokenに一個目
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != tokenizer.EOF {
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

func (p *Parser) peekError(tt tokenizer.TokenType) {
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case tokenizer.AT:
		return p.parseAllocationStatement()
	case tokenizer.LPAREN:
		return p.parseAddressTaggingStatement()
	default:
		return p.parseOpsAndJumpStatement()
	}
}

func (p *Parser) parseAllocationStatement() ast.Statement {
	// be in [@]val
	statement := &ast.AllocationStatement{}

	if !p.peekTokenIs(tokenizer.IDENT) && !p.peekTokenIs(tokenizer.INT) {
		msg := fmt.Sprintf("current token is: %q, expected next token is: %q or %q, got %q(%q)",
			p.currentToken.Literal, tokenizer.IDENT, tokenizer.INT, p.peekToken.Type, p.peekToken.Literal)
		p.addError(msg)
		return nil
	}

	// move to @[val]
	p.nextToken()

	statement.Value = p.currentToken
	statement.Line = p.currentLine
	p.currentLine += 1

	// move to @val[]
	p.nextToken()

	return statement
}

func (p *Parser) expectPeek(tt tokenizer.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	} else {
		p.peekError(tt)
		return false
	}
}

func (p *Parser) peekTokenIs(tt tokenizer.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) parseAddressTaggingStatement() ast.Statement {
	// be in [(]val)
	statement := &ast.AddressTaggingStatement{}

	// move to ([val])
	if !p.expectPeek(tokenizer.IDENT) {
		p.addError(fmt.Sprintf("unexpected token. expected=%q, actual=%q", tokenizer.IDENT, p.currentToken.Type))
		for p.currentToken.Type != tokenizer.RPAREN {
			p.nextToken()
		}
		p.nextToken()
		return nil
	}

	statement.Value = p.currentToken
	statement.Line = p.currentLine
	//p.currentLine += 1

	// move to (val[)]
	if !p.expectPeek(tokenizer.RPAREN) {
		return nil
	}
	// move to (val)[]
	p.nextToken()

	return statement
}

// TODO: try to parse expressions
func (p *Parser) parseOpsAndJumpStatement() ast.Statement {
	// be in [D]=M+1;JMP
	statement := &ast.OpsAndJumpStatement{}

	destTokens := []tokenizer.Token{}
	compTokens := []tokenizer.Token{}
	var jumpToken *tokenizer.Token

	// move to D[=]M+1;JMP
	// move to D[;]JMP
	for p.currentToken.Type != tokenizer.ASSIGN && p.currentToken.Type != tokenizer.SEMICOLON {
		destTokens = append(destTokens, p.currentToken)
		p.nextToken()
	}

	if p.currentToken.Type == tokenizer.ASSIGN {
		// be in D[=]M+1;JMP
		// move to D=[M]+1;JMP
		p.nextToken()
		for p.currentToken.Type != tokenizer.SEMICOLON {
			compTokens = append(compTokens, p.currentToken)
			p.nextToken()
		}
		// be in D=M+1[;]JMP
		// move to D=M+1;[JMP]
		p.nextToken()
	} else {
		// be in D[;]JMP
		// move to D;[JMP]
		p.nextToken()
		// no dest part
		for _, token := range destTokens {
			compTokens = append(compTokens, token)
		}
		destTokens = []tokenizer.Token{}
	}

	if p.isJumpToken() {
		jumpToken = &tokenizer.Token{Type: p.currentToken.Type, Literal: p.currentToken.Literal}
		// move to D;JMP[]
		p.nextToken()
	}

	statement.Dest = destTokens
	statement.Comp = compTokens
	statement.Jump = jumpToken
	statement.Line = p.currentLine
	p.currentLine += 1

	return statement
}

func (p *Parser) isJumpToken() bool {
	return p.currentToken.Type == tokenizer.JGT ||
		p.currentToken.Type == tokenizer.JMP ||
		p.currentToken.Type == tokenizer.JEQ ||
		p.currentToken.Type == tokenizer.JGE ||
		p.currentToken.Type == tokenizer.JLE ||
		p.currentToken.Type == tokenizer.JLT ||
		p.currentToken.Type == tokenizer.JNE
}
