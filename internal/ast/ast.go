package ast

import (
	"strings"

	"github.com/Aden-Q/monkey/internal/token"
)

// interface compliance check
var _ Node = (*Program)(nil)
var _ Expression = (*Identifier)(nil)
var _ Expression = (*Integer)(nil)
var _ Statement = (*LetStatement)(nil)
var _ Statement = (*ReturnStatement)(nil)
var _ Statement = (*ExpressionStatement)(nil)

// Node is a common interface for nodes in AST
type Node interface {
	TokenLiteral() string
	// for debug purpose only
	String() string
}

// Expression is a node that produces a value
type Expression interface {
	Node
	expressionNode()
}

// Statement is a node that does not produce a value
type Statement interface {
	Node
	statementNode()
}

// Program is a representation of the AST. It implements the Node interface (root node of AST)
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

// NewProgram creates a program node
func NewProgram(statements ...Statement) *Program {
	return &Program{
		Statements: statements,
	}
}

// -------------- Expressions -------------------

// Identifier implements the Expression interface
type Identifier struct {
	// the identifier token
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// NewIdentifier creates an identifier expression node
func NewIdentifier(literal string) *Identifier {
	return &Identifier{
		Token: token.New(token.IDENT, literal),
		Value: literal,
	}
}

// Integer implements the Expression interface
type Integer struct {
	// the integer token
	Token token.Token
	Value int64
}

func (i *Integer) expressionNode() {}

func (i *Integer) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Integer) String() string {
	return i.Token.Literal
}

// NewInteger creates an integer expression node
func NewInteger(literal string, value int64) *Integer {
	return &Integer{
		Token: token.New(token.INT, literal),
		Value: value,
	}
}

// PrefixExpression implements the Expression interface
// a prefix expression consists of a prefix (-/!) and an operator
type PrefixExpression struct {
	// a token representation of the prefix operator
	Token token.Token
	// the string literal of the prefix operator
	Operator string
	// the expression following the prefix operator
	Operand Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	builder := strings.Builder{}

	builder.WriteString("(")
	builder.WriteString(pe.Operator)
	builder.WriteString(pe.Operand.String())
	builder.WriteString(")")

	return builder.String()
}

// NewPrefixExpression creates a prefix expression node
func NewPrefixExpression(literal string, operand Expression) *PrefixExpression {
	return &PrefixExpression{
		Token:    token.New(token.LookupTokenType(literal), literal),
		Operator: literal,
		Operand:  operand,
	}
}

// InfixExpression implements the Expression interface
type InfixExpression struct {
	// a token representation of the infix operator
	Token token.Token
	// the string literal of the infix operator
	Operator string
	// the expression to the left of the infix expression
	LeftOperand Expression
	// the expression to the right of the infix expression
	RightOperand Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	builder := strings.Builder{}

	builder.WriteString("(")
	builder.WriteString(ie.LeftOperand.String())
	builder.WriteString(ie.Operator)
	builder.WriteString(ie.RightOperand.String())
	builder.WriteString(")")

	return builder.String()
}

// NewInExpression creates an infix expression node
func NewInfixExpression(literal string, leftOperand, rightOperand Expression) *InfixExpression {
	return &InfixExpression{
		Token:        token.New(token.LookupTokenType(literal), literal),
		Operator:     literal,
		LeftOperand:  leftOperand,
		RightOperand: rightOperand,
	}
}

// -------------- Statements -------------------

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

// NewLetStatement creates a let statement node
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

// NewReturnStatement creates a return statement node
func NewReturnStatement(value Expression) *ReturnStatement {
	return &ReturnStatement{
		Token: token.New(token.RETURN, "return"),
		Value: value,
	}
}

// ExpressionStatement represents a statement consisting of only one expression
type ExpressionStatement struct {
	Statement
	// the expression
	Expression Expression
}

// NewExpressionStatement creates an expression statement node
func NewExpressionStatement(exp Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expression: exp,
	}
}
