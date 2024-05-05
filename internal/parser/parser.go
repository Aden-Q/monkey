package parser

import (
	"strconv"

	"github.com/aden-q/monkey/internal/token"

	"github.com/aden-q/monkey/internal/ast"
	"github.com/aden-q/monkey/internal/lexer"
)

// interface compliance check
var _ Parser = (*parser)(nil)

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
	// handler for identifier expression
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	// handler for integer expression
	p.registerPrefixParseFn(token.INT, p.parseInteger)
	// handler for boolean expression
	p.registerPrefixParseFn(token.TRUE, p.parseBoolean)
	p.registerPrefixParseFn(token.FALSE, p.parseBoolean)
	// handler for grouped expression
	p.registerPrefixParseFn(token.LPAREN, p.parseGroupedExpression)
	// handler for if expression
	p.registerPrefixParseFn(token.IF, p.parseIfExpression)
	// handler for !something expression
	p.registerPrefixParseFn(token.BANG, p.parsePrefixExpression)
	// handler for -something expression
	p.registerPrefixParseFn(token.MINUS, p.parsePrefixExpression)

	// register infix parse functions
	p.registerInfixParseFn(token.PLUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.MINUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParseFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixParseFn(token.GT, p.parseInfixExpression)
	p.registerInfixParseFn(token.GTE, p.parseInfixExpression)
	p.registerInfixParseFn(token.LT, p.parseInfixExpression)
	p.registerInfixParseFn(token.LTE, p.parseInfixExpression)
	p.registerInfixParseFn(token.EQ, p.parseInfixExpression)
	p.registerInfixParseFn(token.NOT_EQ, p.parseInfixExpression)

	return p
}

func (p *parser) registerPrefixParseFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *parser) registerInfixParseFn(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// TODO: implement it to parse the whole program into a AST
func (p *parser) ParseProgram(text string) (*ast.Program, []error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	errs := []error{}

	// reset the lexer state
	_ = p.l.Read(text)

	// reset the parser state
	// the main reason of doing this is skipping any leading white space/newline char
	// we need to do nextToken twice to populate both the current token and the next token
	p.curToken = token.Token{}
	p.peekToken = token.Token{}
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

	if err != nil {
		// this loop is needed when we fail to parse the current statement, skip the rest of it
		for !p.curTokenTypeIs(token.SEMICOLON) && !p.curTokenTypeIs(token.EOF) {
			p.nextToken()
		}

		return stmt, err
	}

	// on successful parsing, we skip the ; token to parse the next potential statement
	p.nextToken()

	return stmt, err
}

// parseLetStatement parses a single let statement
func (p *parser) parseLetStatement() (ast.Statement, error) {
	// expect the next token type to be IDENT
	if !p.peekTokenTypeIs(token.IDENT) {
		// fail to parse this let statement
		return nil, ErrUnexpectedTokenType
	}

	// move forward
	p.nextToken()

	// expect the next token type to be IDENT
	if !p.peekTokenTypeIs(token.ASSIGN) {
		// fail to parse this let statement
		return nil, ErrUnexpectedTokenType
	}

	tok := p.curToken

	// move forward to make p.curToekn be the first token of the expression
	p.nextToken()
	p.nextToken()

	value, err := p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	return ast.NewLetStatement(ast.NewIdentifier(tok.Literal), value), nil
}

// parseReturnStatement parses a single return statement
func (p *parser) parseReturnStatement() (ast.Statement, error) {
	// move forward to make p.curToekn be the first token of the expression
	p.nextToken()

	exp, err := p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	return ast.NewReturnStatement(exp), nil
}

// parseExpressionStatement parses a single expression statement
func (p *parser) parseExpressionStatement() (ast.Statement, error) {
	exp, err := p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	return ast.NewExpressionStatement(exp), err
}

// parseExpression parses a single expression, p.curToken points to the first token of the expression
func (p *parser) parseExpression(precedence int) (ast.Expression, error) {
	prefixFn, ok := p.prefixParseFns[p.curToken.Type]
	if !ok {
		return nil, ErrPrefixParseFnNotFound
	}

	// the prefix expression
	exp, err := prefixFn()
	if err != nil {
		return nil, err
	}

	// recursively parse the remaining part
	for !p.peekTokenTypeIs(token.SEMICOLON) && precedence < token.GetPrecedence(p.peekToken.Type) {
		infixFn, ok := p.infixParseFns[p.peekToken.Type]
		if !ok {
			return exp, ErrInfixParseFnNotFound
		}

		// move forward to make p.curToekn point to the operator of the infix expression
		p.nextToken()

		exp, err = infixFn(exp)
		if err != nil {
			return nil, err
		}
	}

	return exp, err
}

func (p *parser) parseIdentifier() (ast.Expression, error) {
	return ast.NewIdentifier(p.curToken.Literal), nil
}

func (p *parser) parseInteger() (ast.Expression, error) {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, err
	}

	return ast.NewInteger(p.curToken.Literal, value), nil
}

func (p *parser) parseBoolean() (ast.Expression, error) {
	value, err := strconv.ParseBool(p.curToken.Literal)
	if err != nil {
		return nil, err
	}

	return ast.NewBoolean(value), nil
}

func (p *parser) parseGroupedExpression() (ast.Expression, error) {
	// move forward to make p.curToken point to the first token after the ( token
	p.nextToken()

	// recursively parse the expression after the prefix token
	exp, err := p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	if !p.peekTokenTypeIs(token.RPAREN) {
		return nil, ErrUnexpectedTokenType
	}

	// move forward to make p.curToken point to ), do not skip ) because we need to check peek token
	// in a infix expression
	p.nextToken()

	return exp, nil
}

func (p *parser) parseBlockStatement() (*ast.BlockStatement, error) {
	// skip the { token pointed by p.curToken
	p.nextToken()

	return nil, nil
}

func (p *parser) parseIfExpression() (ast.Expression, error) {
	if !p.peekTokenTypeIs(token.LPAREN) {
		return nil, ErrUnexpectedTokenType
	}

	// move forward to make p.curToken point to ( so that we can parse the grouped expressions
	p.nextToken()

	condition, err := p.parseGroupedExpression()
	if err != nil {
		return nil, err
	}

	if !p.peekTokenTypeIs(token.LBRACE) {
		return nil, ErrUnexpectedTokenType
	}

	// move forward so that p.curToken point to {
	p.nextToken()

	consequence, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	return ast.NewIfExpression(condition, consequence, nil), nil
}

func (p *parser) parsePrefixExpression() (ast.Expression, error) {
	prefixToken := p.curToken

	// move forward to make p.curToekn points to the operand expression
	p.nextToken()

	// recursively parse the expression after the prefix token
	operand, err := p.parseExpression(token.PREFIX)
	if err != nil {
		return nil, err
	}

	return ast.NewPrefixExpression(prefixToken.Literal, operand), nil
}

func (p *parser) parseInfixExpression(leftOperand ast.Expression) (ast.Expression, error) {
	operatorToken := p.curToken
	precedence := token.GetPrecedence(operatorToken.Type)

	// move forward to make p.curToekn points to the right operand expression
	p.nextToken()

	rightOperand, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}

	return ast.NewInfixExpression(operatorToken.Literal, leftOperand, rightOperand), nil
}

// nextToken uses the lexer to read the next token and mutate the parser's state
func (p *parser) nextToken() {
	tok := p.l.NextToken()

	p.curToken = p.peekToken
	p.peekToken = tok
}

// peekTokenTypeIs examines whether the current token type is the expected one
func (p *parser) curTokenTypeIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

// peekTokenTypeIs examines whether the peek token type is the expected one
func (p *parser) peekTokenTypeIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}
