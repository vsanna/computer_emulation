package vm_tokenizer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) IsZero() bool {
	return t.Type == ZERO_VAL
}

const (
	ZERO_VAL = "ZERO_VAL" // ゼロ値用
	IDENT    = "IDENT"
	ILLEGAL  = "ILLEGAL"

	EOF = "EOF"

	INT = "INT"

	PUSH = "PUSH"
	//POP  = "POP"

	ADD = "ADD"
	//SUB = "SUB"
	//
	//NEQ = "NEQ"
	//EQ = "EQ"
	//GT = "GT"
	//LT = "LT"
	//AND = "AND"
	//OR = "OR"
	//NOT = "NOT"

	CONSTANT = "CONSTANT"
	//LOCAL = "LOCAL"
	//ARGUMENT = "ARGUMENT"
	//THIS = "THIS"
	//THAT = "THAT"
	//POINTER = "POINTER"
	//TEMP = "TEMP"
	//STATIC = "STATIC"
)

var keywords = map[string]TokenType{
	"push":     PUSH,
	"add":      ADD,
	"constant": CONSTANT,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}

func ZeroToken() Token {
	return Token{Literal: "", Type: ZERO_VAL}
}
