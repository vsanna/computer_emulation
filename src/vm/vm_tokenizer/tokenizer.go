package vm_tokenizer

import "strings"

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Tokenizer {
	// remove comments
	inputWithoutComments := withoutComment(input)
	t := &Tokenizer{input: inputWithoutComments}
	t.readChar()
	return t
}

func withoutComment(input string) string {
	result := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}

		lineWithoutComment := strings.Split(trimmedLine, "//")[0]

		if len(lineWithoutComment) == 0 {
			continue
		}
		result += lineWithoutComment + "\n"
	}
	return result
}

func (l *Tokenizer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // null文字
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Tokenizer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.ch {
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	// 1byteのtokenのみの前提.
	l.readChar()
	return tok
}

// pos = 0のとき、次のcharをpeek
func (l *Tokenizer) peekChar(pos int) byte {
	if l.readPosition+pos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition+pos]
	}
}

func (l *Tokenizer) readIdentifier() string {
	position := l.position
	// 英字のみからなる文字列を読み進める
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Tokenizer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?' || ch == '!'
}

func (l *Tokenizer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
