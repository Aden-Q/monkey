package parser

import (
	"github.com/Aden-Q/monkey/internal/token"

	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
)

var _ Parser = (*parser)(nil)

// a Pratt Parser interface
type Parser interface {
	// TODO: add interface methods
	ParseProgram(text string) *ast.Program
}

// a Pratt Parser implementation
type parser struct {
	l lexer.Lexer

	// curtoken and peekToken are used to keep track of
	// the current parsing progress, the object is stateful
	curToken  token.Token
	peekToken token.Token
}

func New(l lexer.Lexer) Parser {
	return &parser{
		l: l,
	}
}

// TODO: implement it to parse the whole program into a AST
func (p *parser) ParseProgram(text string) *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	_ = p.l.Read(text)

	for p.curToken.Type != token.EOF {
		// parse a single statement every time
		stmt := p.parseStatment()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

// parseStatment parses a single statement
func (p *parser) parseStatment() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}

	return nil
}

// parseLetStatement parses a single let statement
func (p *parser) parseLetStatement() ast.Statement {
	// expect the next token type to be IDENT
	if p.peekToken.Type != token.IDENT {
		// fail to parse this let statement
		return nil
	}

	// move forward
	p.nextToken()

	// expect the next token type to be IDENT
	if p.peekToken.Type != token.ASSIGN {
		// fail to parse this let statement
		return nil
	}

	stmt := ast.NewLetStatement(&ast.Identifier{
		Token: p.curToken,
	}, nil)

	for !(p.curToken.Type == token.SEMICOLON) && !(p.curToken.Type == token.EOF) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement parses a single return statement
func (p *parser) parseReturnStatement() ast.Statement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	return stmt
}

// nextToken uses the lexer to read the next token and mutate the parser's state
func (p *parser) nextToken() {
	tok := p.l.NextToken()

	p.curToken = p.peekToken
	p.peekToken = tok
}
