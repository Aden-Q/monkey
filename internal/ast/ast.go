package ast

import (
	"strings"

	"github.com/Aden-Q/monkey/internal/token"
)

// interface compliance check
var _ Node = (*Program)(nil)
var _ Expression = (*Identifier)(nil)
var _ Statement = (*LetStatement)(nil)
var _ Statement = (*ReturnStatement)(nil)
var _ Statement = (*ExpressionStatement)(nil)

// Node is a common interface for nodes in AST
type Node interface {
	TokenLiteral() string
	// for debug purpose only
	String() string
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

func (p *Program) String() string {
	builder := strings.Builder{}

	for _, s := range p.Statements {
		builder.WriteString(s.String())
	}

	return builder.String()
}

// Identifier implements the Expression interface because an identifier object
// can the right value of a statement, meaning that it can evaluate to some value, after it's assgined
type Identifier struct {
	// the identifier token
	Token token.Token
	// associated value of this identifier, represented as a string for now
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func NewIdentifier(literal string, value string) *Identifier {
	return &Identifier{
		Token: token.New(token.IDENT, literal),
		Value: value,
	}
}

// LetStatement represents the let statement
type LetStatement struct {
	// the let token
	Token token.Token
	// the identifier
	Identifier *Identifier
	// the expression value on the right side of the statement
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	builder := strings.Builder{}

	builder.WriteString(ls.TokenLiteral() + " ")
	builder.WriteString(ls.Identifier.TokenLiteral() + " ")
	builder.WriteString("= ")

	if ls.Value != nil {
		builder.WriteString(ls.Value.String())
	}

	builder.WriteString(";")

	return builder.String()
}

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

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	builder := strings.Builder{}

	builder.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		builder.WriteString(rs.Value.String())
	}

	builder.WriteString(";")

	return builder.String()
}

func NewReturnStatement(value Expression) *ReturnStatement {
	return &ReturnStatement{
		Token: token.New(token.RETURN, "return"),
		Value: value,
	}
}

// ExpressionStatement represents a statement consisting of only one expression
type ExpressionStatement struct {
	// the first token of the expression
	Token token.Token
	// the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression == nil {
		return ""
	}

	return es.Expression.String()
}

func NewExpressionStatement(exp Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Token:      token.New(token.RETURN, "return"),
		Expression: exp,
	}
}
