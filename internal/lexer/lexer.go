package lexer

import (
	"github.com/Aden-Q/monkey/internal/token"
)

type Lexer struct {
	input    string
	position uint32 // current position index in input
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	return l
}

func (l *Lexer) NextToken() (token.Token, bool) {
	if !l.hasNext() {
		return token.Token{}, false
	}

	l.skipWhiteSpaces()

	var tok token.Token
	ok := true

	ch := l.input[l.position]

	switch ch {
	// operators
	case '=', '+', '-', '!', '*', '/', '<', '>':
		fallthrough
	// delimiters
	case ',', ';', '(', ')', '{', '}':
		ch := l.readChar()
		tok = token.New(token.LookupTokenType(ch), ch)
	default:
		// read identifier
		if isLetter(ch) {
			literal := l.readWord()
			tok = token.New(token.LookupTokenType(literal), literal)
			ok = true
		} else if isDigit(ch) {
			literal := l.readInt()
			tok = token.New(token.LookupTokenType(literal), literal)
			ok = true
		} else {
			tok = token.New(token.ILLEGAL, string(ch))
			ok = false
		}
	}

	return tok, ok
}

func (l *Lexer) hasNext() bool {
	return l.position < uint32(len(l.input))
}

func (l *Lexer) readChar() string {
	if !l.hasNext() {
		return ""
	}

	l.position++

	return l.input[l.position-1 : l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// read a word starting from the current position, and move the offset forward
func (l *Lexer) readWord() string {
	startPos := l.position

	for {
		if !l.hasNext() {
			break
		}

		ch := l.input[l.position]
		if !isLetter(ch) {
			break
		}

		l.position++
	}

	return l.input[startPos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readInt() string {
	startPos := l.position

	for {
		if !l.hasNext() {
			break
		}

		ch := l.input[l.position]
		if !isDigit(ch) {
			break
		}

		l.position++
	}

	return l.input[startPos:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for {
		if !l.hasNext() {
			break
		}

		ch := l.input[l.position]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			l.position += 1
		} else {
			break
		}
	}
}
