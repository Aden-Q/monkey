package lexer

import (
	"github.com/Aden-Q/monkey/internal/token"
)

type Lexer struct {
	input    string
	position uint32 // current position index in input
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) NextToken() (*token.Token, bool) {
	if !l.hasNext() {
		return nil, false
	}

	var tok *token.Token
	ok := true

	switch l.input[l.position] {
	case '=':
		tok = token.New(token.ASSIGN, "=")
	case '+':
		tok = token.New(token.PLUS, "+")
	case ',':
		tok = token.New(token.COMMA, ",")
	case ';':
		tok = token.New(token.SEMICOLON, ";")
	case '(':
		tok = token.New(token.LPAREN, "(")
	case ')':
		tok = token.New(token.RPAREN, ")")
	case '{':
		tok = token.New(token.LBRACE, "{")
	case '}':
		tok = token.New(token.RBRACE, "}")
	default:
		tok = token.New(token.ILLEGAL, "ILLEGAL")
		ok = false
	}

	l.position++

	return tok, ok
}

func (l *Lexer) hasNext() bool {
	return l.position < uint32(len(l.input))
}
