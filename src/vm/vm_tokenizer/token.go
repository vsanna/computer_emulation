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
	POP  = "POP"

	ADD = "ADD"
	SUB = "SUB"

	EQ  = "EQ"
	GT  = "GT"
	LT  = "LT"
	AND = "AND"
	OR  = "OR"

	NOT = "NOT"
	NEG = "NEG"

	CONSTANT = "CONSTANT"
	// dynamic
	LOCAL    = "LOCAL"
	ARGUMENT = "ARGUMENT"
	THIS     = "THIS"
	THAT     = "THAT"
	POINTER  = "POINTER"
	// static
	TEMP   = "TEMP"
	STATIC = "STATIC"
)

var keywords = map[string]TokenType{
	"push":     PUSH,
	"pop":      POP,
	"add":      ADD,
	"sub":      SUB,
	"and":      AND,
	"or":       OR,
	"not":      NOT,
	"neg":      NEG,
	"eq":       EQ,
	"gt":       GT,
	"lt":       LT,
	"constant": CONSTANT,
	"argument": ARGUMENT,
	"local":    LOCAL,
	"this":     THIS,
	"that":     THAT,
	"temp":     TEMP,
	"static":   STATIC,
	"pointer":  POINTER,
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
