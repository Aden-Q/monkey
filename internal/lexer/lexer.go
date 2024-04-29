package lexer

import (
	"github.com/Aden-Q/monkey/internal/token"
)

type Lexer struct {
	input    string
	position uint32 // current position index in input
}

func New() *Lexer {
	l := &Lexer{}

	return l
}

func (l *Lexer) Read(input string) int {
	l.input = input
	l.position = 0

	return len(input)
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
	// operators with two characters
	case '=', '!':
		if l.peekNextNextChar() == '=' {
			ch := l.readChar() + l.readChar()
			tok = token.New(token.LookupTokenType(ch), ch)
		} else {
			ch := l.readChar()
			tok = token.New(token.LookupTokenType(ch), ch)
		}
	// operators with a single character
	case '+', '-', '*', '/', '<', '>':
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

// hasNext checks whether there are characters remaining
func (l *Lexer) hasNext() bool {
	return l.position < uint32(len(l.input))
}

// peekNextNextChar looks at the next character after the next character
func (l *Lexer) peekNextNextChar() byte {
	if l.position+1 > uint32(len(l.input))-1 {
		return 0
	}

	return l.input[l.position+1]
}

// readChar reads a single char at the current offset and move the ptr forward by 1
func (l *Lexer) readChar() string {
	if !l.hasNext() {
		return ""
	}

	l.position++

	return l.input[l.position-1 : l.position]
}

// isLetter check whether a character is allow be to in an identifier
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

// isLetter check whether a character is an digit
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

// skipWhiteSpaces skips all white spaces starting at the current position, including newline characters
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
