package ast

import "github.com/Aden-Q/monkey/internal/token"

// interface compliance check
var _ Node = (*Program)(nil)
var _ Statement = (*statement)(nil)
var _ Statement = (*LetStatement)(nil)

// Node is a common interface for nodes in AST
type Node interface {
	TokenLiteral() string
}

// Statement node is a node that does not produce a value
type Statement interface {
	Node
	statementNode()
}

func (s *statement) statementNode() {}

func (s *statement) TokenLiteral() string {
	return s.token.Literal
}

// Expression node is a node that produces a value
type Expression interface {
	Node
	expressionNode()
}

type statement struct {
	token token.Token
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
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement represents the let statement
type LetStatement struct {
	statement
	Name  *Identifier
	Value Expression
}

// ReturnStatement represents the return statement
type ReturnStatement struct {
	statement
	Value Expression
}
