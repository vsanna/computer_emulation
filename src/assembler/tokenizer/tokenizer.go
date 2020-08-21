package tokenizer

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Tokenizer {
	t := &Tokenizer{input: input}
	t.readChar()
	return t
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
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '@':
		tok = newToken(AT, l.ch)
	case '!':
		tok = newToken(BANG, l.ch)
	case '&':
		tok = newToken(AND, l.ch)
	case '|':
		tok = newToken(OR, l.ch)
	case 'A':
		if l.peekChar(0) == 'D' {
			if l.peekChar(1) == 'M' {
				if isLetter(l.peekChar(2)) {
					// if Mの次が英字 -> ident
					tok.Literal = l.readIdentifier()
					tok.Type = LookupIdent(tok.Literal)
					return tok
				} else {
					// その他記号 -> A_REG
					tok = newToken(A_REG, l.ch)
				}
			} else {
				if isLetter(l.peekChar(1)) {
					// if Dの次がM以外の英字 -> ident
					tok.Literal = l.readIdentifier()
					tok.Type = LookupIdent(tok.Literal)
					return tok
				} else {
					// その他記号 -> A_REG
					tok = newToken(A_REG, l.ch)
				}
			}
		} else if l.peekChar(0) == 'M' {
			if isLetter(l.peekChar(1)) {
				// if Mの次が英字 -> ident
				tok.Literal = l.readIdentifier()
				tok.Type = LookupIdent(tok.Literal)
				return tok
			} else {
				// その他記号 -> A_REG
				tok = newToken(A_REG, l.ch)
			}
		} else {
			if isLetter(l.peekChar(1)) {
				// if Aの次がDでもMでもない英字 -> ident
				tok.Literal = l.readIdentifier()
				tok.Type = LookupIdent(tok.Literal)
				return tok
			} else {
				// その他記号 -> A_REG
				tok = newToken(A_REG, l.ch)
			}
		}
	case 'D':
		if l.peekChar(0) == 'M' {
			if isLetter(l.peekChar(1)) {
				// if Mの次が英字 -> ident
				tok.Literal = l.readIdentifier()
				tok.Type = LookupIdent(tok.Literal)
				return tok
			} else {
				// その他記号 -> D_REG
				tok = newToken(D_REG, l.ch)
			}
		} else {
			if isLetter(l.peekChar(0)) {
				// if Dの次がMでない英字 -> ident
				tok.Literal = l.readIdentifier()
				tok.Type = LookupIdent(tok.Literal)
				return tok
			} else {
				// その他記号 -> D_REG
				tok = newToken(D_REG, l.ch)
			}
		}
	case 'M':
		if isLetter(l.peekChar(0)) {
			// if Mの次が英字 -> ident
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else {
			// その他記号 -> D_REG
			tok = newToken(MEMORY, l.ch)
		}
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
	// 英数字のみからなる文字列を読み進める
	for isLetter(l.ch) || isDigit(l.ch) {
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
