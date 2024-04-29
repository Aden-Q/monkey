package parser

import (
	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
)

var _ Parser = (*parser)(nil)

type Parser interface {
	// TODO: add interface methods
	ParseProgram() *ast.Program
}

type parser struct {
	l lexer.Lexer
}

func New(l lexer.Lexer) Parser {
	return &parser{
		l: l,
	}
}

// TODO: implement it to parse the whole program into a AST
func (p *parser) ParseProgram() *ast.Program {
	return nil
}
