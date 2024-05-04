package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Aden-Q/monkey/internal/token"

	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
)

const (
	LOWEST      = iota
	EQUALS      // ==
	LESSGREATER // >, >=, <, <=
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // fn(X)
)

// interface compliance check
var _ Parser = (*parser)(nil)

var (
	ErrUnexpectedTokenType = errors.New("unexpected token type")
)

type (
	prefixParseFn func() (ast.Expression, error)
	infixParseFn  func(ast.Expression) (ast.Expression, error)
)

// a Pratt Parser interface
type Parser interface {
	// TODO: add more interface methods here
	ParseProgram(text string) (*ast.Program, []error)
}

// a Pratt Parser implementation
type parser struct {
	l lexer.Lexer

	// curtoken and peekToken are used to keep track of
	// the current parsing progress, the object is stateful
	curToken  token.Token
	peekToken token.Token

	// parse functions for expressions
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l lexer.Lexer) Parser {
	p := &parser{
		l:              l,
		prefixParseFns: make(map[token.TokenType]prefixParseFn, 0),
		infixParseFns:  make(map[token.TokenType]infixParseFn, 0),
	}

	// register prefix parse functions
	p.registerPrefixFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixFn(token.INT, p.parseInteger)

	return p
}

func (p *parser) registerPrefixFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
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

	// in each iteration of the for loop here, we parse a single statement separated by a semicolon ;
	for p.curToken.Type != token.EOF {
		// parse a single statement every time
		stmt, err := p.parseStatment()
		if err != nil {
			errs = append(errs, err)
		}

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		// move to the start of the next statement
		p.nextToken()
	}

	return program, errs
}

// parseStatment parses a single statement
func (p *parser) parseStatment() (ast.Statement, error) {
	// TODO: check how to propagate errors when the current token is not a statement indicator
	// make sure not to produce duplicate errors for the same statement
	var stmt ast.Statement
	var err error

	switch p.curToken.Type {
	case token.LET:
		stmt, err = p.parseLetStatement()
	case token.RETURN:
		stmt, err = p.parseReturnStatement()
	default:
		stmt, err = p.parseExpressionStatement()
	}

	// move to the end of the current statement, indicated by a semicolon;
	for !(p.curToken.Type == token.SEMICOLON) && !(p.curToken.Type == token.EOF) {
		p.nextToken()
	}

	return stmt, err
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

	tok := p.curToken

	// move forward to make p.curToekn be the first token of the expression
	p.nextToken()
	p.nextToken()

	// TODO: parse the expression after the return token
	exp, _ := p.parseExpression(LOWEST)

	return ast.NewLetStatement(&ast.Identifier{
		Token: tok,
	}, exp), nil
}

// parseReturnStatement parses a single return statement
func (p *parser) parseReturnStatement() (ast.Statement, error) {
	// move forward to make p.curToekn be the first token of the expression
	p.nextToken()

	// TODO: parse the expression after the return token
	exp, _ := p.parseExpression(LOWEST)

	return ast.NewReturnStatement(exp), nil
}

// parseExpressionStatement parses a single expression statement
func (p *parser) parseExpressionStatement() (ast.Statement, error) {
	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	return ast.NewExpressionStatement(exp), err
}

// parseExpression parses a single expression, p.curToken points to the first token of the expression
func (p *parser) parseExpression(precedence int) (ast.Expression, error) {
	prefixFn, ok := p.prefixParseFns[p.curToken.Type]
	if !ok {
		return nil, nil
	}

	return prefixFn()
}

func (p *parser) parseIdentifier() (ast.Expression, error) {
	return ast.NewIdentifier(p.curToken.Literal), nil
}

func (p *parser) parseInteger() (ast.Expression, error) {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("parse integer %s: %v", p.curToken.Literal, err)
	}

	return ast.NewInteger(p.curToken.Literal, value), nil
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
