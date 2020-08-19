package tokenizer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	IDENT   = "IDENT"
	ILLEGAL = "ILLEGAL"

	EOF    = "EOF"
	AT     = "@"
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	OR   = "|"
	AND  = "&"
	BANG = "!"

	LPAREN    = "("
	RPAREN    = ")"
	SEMICOLON = ";"

	MEMORY = "MEMORY"
	A_REG  = "A_REG"
	D_REG  = "D_REG"

	INT = "INT"

	JGT = "JGT"
	JEQ = "JEQ"
	JGE = "JGE"
	JLT = "JLT"
	JNE = "JNE"
	JLE = "JLE"
	JMP = "JMP"
)

var keywords = map[string]TokenType{
	"M":   MEMORY,
	"A":   A_REG,
	"D":   D_REG,
	"JGT": JGT,
	"JEQ": JEQ,
	"JGE": JGE,
	"JLT": JLT,
	"JNE": JNE,
	"JLE": JLE,
	"JMP": JMP,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
