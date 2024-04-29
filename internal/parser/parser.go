package parser

import "github.com/Aden-Q/monkey/internal/lexer"

type Parser interface {
}

type parser struct {
	l lexer.Lexer
}
