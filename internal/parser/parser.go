package parser

import (
	"errors"

	"github.com/Aden-Q/monkey/internal/token"

	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
)

// interface compliance check
var _ Parser = (*parser)(nil)

var (
	ErrUnexpectedTokenType = errors.New("unexpected token type")
)

// a Pratt Parser interface
type Parser interface {
	// TODO: add interface methods
	ParseProgram(text string) (*ast.Program, []error)
}

// a Pratt Parser implementation
type parser struct {
	l lexer.Lexer

	// curtoken and peekToken are used to keep track of
	// the current parsing progress, the object is stateful
	curToken  token.Token
	peekToken token.Token

	// parsing errors

}

func New(l lexer.Lexer) Parser {
	return &parser{
		l: l,
	}
}

// TODO: implement it to parse the whole program into a AST
func (p *parser) ParseProgram(text string) (*ast.Program, []error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	errs := []error{}

	_ = p.l.Read(text)
	// the main reason of doing this is skipping any leading white space/newline char
	// we need to do nextToken twice to populate both the current token and the next token
	p.nextToken()
	p.nextToken()

	for p.curToken.Type != token.EOF {
		// parse a single statement every time
		stmt, err := p.parseStatment()
		if err != nil {
			errs = append(errs, err)
		}

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program, errs
}

// parseStatment parses a single statement
func (p *parser) parseStatment() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}

	// TODO: check how to propagate errors when the current token is not a statement indicator
	// make sure not to produce duplicate errors for the same statement
	return nil, nil
}

// parseLetStatement parses a single let statement
func (p *parser) parseLetStatement() (ast.Statement, error) {
	// expect the next token type to be IDENT
	if !p.expectPeekTokenType(token.IDENT) {
		// fail to parse this let statement
		return nil, ErrUnexpectedTokenType
	}

	// move forward
	p.nextToken()

	// expect the next token type to be IDENT
	if !p.expectPeekTokenType(token.ASSIGN) {
		// fail to parse this let statement
		return nil, ErrUnexpectedTokenType
	}

	stmt := ast.NewLetStatement(&ast.Identifier{
		Token: p.curToken,
	}, nil)

	// TODO: parse the expression after the = token
	for !(p.curToken.Type == token.SEMICOLON) && !(p.curToken.Type == token.EOF) {
		p.nextToken()
	}

	return stmt, nil
}

// parseReturnStatement parses a single return statement
func (p *parser) parseReturnStatement() (ast.Statement, error) {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	return stmt, nil
}

// nextToken uses the lexer to read the next token and mutate the parser's state
func (p *parser) nextToken() {
	tok := p.l.NextToken()

	p.curToken = p.peekToken
	p.peekToken = tok
}

// expectPeekTokenType examines whether the peek token type is the expected one
func (p *parser) expectPeekTokenType(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}
