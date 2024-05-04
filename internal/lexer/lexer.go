package lexer

import (
	"github.com/Aden-Q/monkey/internal/token"
)

var _ Lexer = (*lexer)(nil)

type Lexer interface {
	// Read reads the input text and stores into the buffer
	Read(text string) int
	// NextToken reads the next token starting at the current offset and move the ptr forward
	NextToken() token.Token
}

type lexer struct {
	buf      string
	position uint32 // current position index in input
}

func New() Lexer {
	return &lexer{}
}

func (l *lexer) Read(text string) int {
	l.buf = text
	l.position = 0

	return len(text)
}

func (l *lexer) NextToken() token.Token {
	l.skipWhiteSpaces()

	if !l.hasNext() {
		return token.Token{
			Type:    token.EOF,
			Literal: "eof",
		}
	}

	var tok token.Token
	ch := l.buf[l.position]

	switch ch {
	// operators with two characters
	case '=', '!', '<', '>':
		if l.peekNextNextChar() == '=' {
			ch := l.readChar() + l.readChar()
			tok = token.New(token.LookupTokenType(ch), ch)
		} else {
			ch := l.readChar()
			tok = token.New(token.LookupTokenType(ch), ch)
		}
	// operators with a single character
	case '+', '-', '*', '/':
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
		} else if isDigit(ch) {
			literal := l.readInt()
			tok = token.New(token.LookupTokenType(literal), literal)
		} else {
			tok = token.New(token.ILLEGAL, string(ch))
		}
	}

	return tok
}

// hasNext checks whether there are characters remaining
func (l *lexer) hasNext() bool {
	return l.position < uint32(len(l.buf))
}

// peekNextNextChar looks at the next character after the next character
func (l *lexer) peekNextNextChar() byte {
	if l.position+1 > uint32(len(l.buf))-1 {
		return 0
	}

	return l.buf[l.position+1]
}

// readChar reads a single char at the current offset and move the ptr forward by 1
func (l *lexer) readChar() string {
	if !l.hasNext() {
		return ""
	}

	l.position++

	return l.buf[l.position-1 : l.position]
}

// isLetter check whether a character is allow be to in an identifier
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// read a word starting from the current position, and move the offset forward
func (l *lexer) readWord() string {
	startPos := l.position

	for {
		if !l.hasNext() {
			break
		}

		ch := l.buf[l.position]
		if !isLetter(ch) {
			break
		}

		l.position++
	}

	return l.buf[startPos:l.position]
}

// isLetter check whether a character is an digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *lexer) readInt() string {
	startPos := l.position

	for {
		if !l.hasNext() {
			break
		}

		ch := l.buf[l.position]
		if !isDigit(ch) {
			break
		}

		l.position++
	}

	return l.buf[startPos:l.position]
}

// skipWhiteSpaces skips all white spaces starting at the current position, including newline characters
func (l *lexer) skipWhiteSpaces() {
	for {
		if !l.hasNext() {
			break
		}

		ch := l.buf[l.position]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			l.position += 1
		} else {
			break
		}
	}
}
