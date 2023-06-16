package lexer

import (
	"github.com/savannahostrowski/gossamer/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EQUAL, l.ch)
		}
	case '(':
		tok = newToken(token.LPAR, l.ch)
	case ')':
		tok = newToken(token.RPAR, l.ch)
	case '[':
		tok = newToken(token.LSQB, l.ch)
	case ']':
		tok = newToken(token.RSQB, l.ch)
	case ';':
		tok = newToken(token.SEMI, l.ch)
	case ':':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.COLONEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.COLON, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUSEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUSEQUAL, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.RARROW, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '/':
		slashtok := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			literal := string(slashtok) + string(l.ch)
			tok = token.Token{Type: token.SLASHEQUAL, Literal: literal}
		} else if l.peekChar() == '/' {
			secondslashtok := l.ch
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				literal := string(slashtok) + string(secondslashtok) + string(l.ch)
				tok = token.Token{Type: token.DOUBLESLASHEQUAL, Literal: literal}
			} else {
				literal := string(secondslashtok) + string(l.ch)
				tok = token.Token{Type: token.DOUBLESLASH, Literal: literal}
			}
		} else {
			tok = newToken(token.SLASH, slashtok)
		}
	case '*':
		startok := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.STAREQUAL, Literal: string(startok) + string(l.ch)}
		} else if l.peekChar() == '*' {
			secondstartok := l.ch
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				literal := string(startok) + string(secondstartok) + string(l.ch)
				tok = token.Token{Type: token.DOUBLESTAREQUAL, Literal: literal}
			} else {
				tok = token.Token{Type: token.DOUBLESTAR, Literal: string(secondstartok) + string(l.ch)}
			}
		} else {
			tok = newToken(token.STAR, startok)
		}
	case '<':
		lesstok := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.LESSEQUAL, Literal: string(lesstok) + string(l.ch)}
		} else if l.peekChar() == '<' {
			secondlesstok := l.ch
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				literal := string(lesstok) + string(secondlesstok) + string(l.ch)
				tok = token.Token{Type: token.LEFTSHIFTEQUAL, Literal: literal}
			} else {
				literal := string(secondlesstok) + string(l.ch)
				tok = token.Token{Type: token.LEFTSHIFT, Literal: literal}
			}
		} else {
			tok = newToken(token.LESS, lesstok)
		}
	case '>':
		greatertok := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.GREATEREQUAL, Literal: string(greatertok) + string(l.ch)}
		} else if l.peekChar() == '>' {
			secondgreatertok := l.ch
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				literal := string(greatertok) + string(secondgreatertok) + string(l.ch)
				tok = token.Token{Type: token.RIGHTSHIFTEQUAL, Literal: literal}
			} else {
				literal := string(secondgreatertok) + string(l.ch)
				tok = token.Token{Type: token.RIGHTSHIFT, Literal: literal}
			}
		} else {
			tok = newToken(token.GREATER, greatertok)
		}
	case '!':
		// !=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOTEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			// !
			tok = newToken(token.EXCLAMATION, l.ch)
		}
	case '~':
		tok = newToken(token.TILDE, l.ch)
	case '^':
		// ^=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.CIRCUMFLEXEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			// ^
			tok = newToken(token.CIRCUMFLEX, l.ch)
		}
	case '.':
		// ...
		if l.peekChar() == '.' {
			ch := l.ch
			l.readChar()
			if l.peekChar() == '.' {
				l.readChar()
				tok = token.Token{Type: token.ELLIPSIS, Literal: string(ch) + string(l.ch) + string(l.ch)}
			} else {
				tok = token.Token{Type: token.ERRORTOKEN, Literal: string(ch) + string(l.ch)}
			}
		} else {
			// .
			tok = newToken(token.DOT, l.ch)
		}
	case '%':
		// %=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PERCENTEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			// %
			tok = newToken(token.PERCENT, l.ch)
		}
	case '@':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ATEQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.AT, l.ch)
		}
	case '&':
		// &=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AMPEREQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			// &
			tok = newToken(token.AMPER, l.ch)
		}
	case '|':
		// |=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.VBAREQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			// |
			tok = newToken(token.VBAR, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.ENDMARKER
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ERRORTOKEN, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// TODO: skipping whitespace for now but need to figure out how to handle this
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
