package ast

import "github.com/Aden-Q/monkey/internal/token"

// interface compliance check
var _ Node = (*Program)(nil)
var _ Expression = (*Identifier)(nil)
var _ Statement = (*LetStatement)(nil)
var _ Statement = (*ReturnStatement)(nil)

// Node is a common interface for nodes in AST
type Node interface {
	TokenLiteral() string
}

// Statement is a node that does not produce a value
type Statement interface {
	Node
	statementNode()
}

// Expression is a node that produces a value
type Expression interface {
	Node
	expressionNode()
}

// Program is a representation of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}

	// the root of the AST is the first node
	return p.Statements[0].TokenLiteral()
}

type Identifier struct {
	// the identifier token
	Token token.Token
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

// LetStatement represents the let statement
type LetStatement struct {
	// the let token
	Token token.Token
	// the identifier
	Identifier *Identifier
	// the expression value on the right side of the statement
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

func NewLetStatement(identifier *Identifier, value Expression) *LetStatement {
	return &LetStatement{
		Token:      token.New(token.LET, "let"),
		Identifier: identifier,
		Value:      value,
	}
}

// ReturnStatement represents the return statement
type ReturnStatement struct {
	// the return token
	Token token.Token
	// the expression value on the right of the return keyword
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func NewReturnStatement(value Expression) *ReturnStatement {
	return &ReturnStatement{
		Token: token.New(token.RETURN, "return"),
		Value: value,
	}
}
