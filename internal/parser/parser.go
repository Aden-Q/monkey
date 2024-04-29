package parser

import "github.com/Aden-Q/monkey/internal/lexer"

var _ Parser = (*parser)(nil)

type Parser interface {
	// TODO: add interface methods
}

type parser struct {
	l lexer.Lexer
}

func New(l lexer.Lexer) Parser {
	return &parser{
		l: l,
	}
}
