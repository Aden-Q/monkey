package ast

import (
	"strings"

	"github.com/aden-q/monkey/internal/token"
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

// NewProgram creates a Program node
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

// NewIdentifier creates an Identifier node
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

// NewInteger creates an Integer node
func NewInteger(literal string, value int64) *Integer {
	return &Integer{
		Token: token.New(token.INT, literal),
		Value: value,
	}
}

// Boolean implements the Expression interface
type Boolean struct {
	// the boolean token
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

// NewBoolean creates a Boolean node
func NewBoolean(value bool) *Boolean {
	var tok token.Token
	if value {
		tok = token.New(token.TRUE, "true")
	} else {
		tok = token.New(token.FALSE, "false")
	}

	return &Boolean{
		Token: tok,
		Value: value,
	}
}

// IfExpression implements the Expression interface
type IfExpression struct {
	// the if token
	Token token.Token
	// the condition expression
	Condition Expression
	// consequence when the condition is true
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	builder := strings.Builder{}

	builder.WriteString("if" + " ")
	builder.WriteString(ie.Condition.String() + " ")
	builder.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		builder.WriteString(" else ")
		builder.WriteString(ie.Alternative.String())
	}

	return builder.String()
}

// NewIfExpression creates an IfExpression node
func NewIfExpression(condition Expression, consequence *BlockStatement, alternative *BlockStatement) *IfExpression {
	return &IfExpression{
		Token:       token.New(token.IF, "if"),
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
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

// NewPrefixExpression creates a PrefixExpression node
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
	builder.WriteString(ie.LeftOperand.String() + " ")
	builder.WriteString(ie.Operator + " ")
	builder.WriteString(ie.RightOperand.String())
	builder.WriteString(")")

	return builder.String()
}

// NewInExpression creates an InfixExpression node
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

// NewLetStatement creates a LetStatement node
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

// NewReturnStatement creates a ReturnStatement node
func NewReturnStatement(value Expression) *ReturnStatement {
	return &ReturnStatement{
		Token: token.New(token.RETURN, "return"),
		Value: value,
	}
}

// ExpressionStatement represents a statement consisting of only one expression
type ExpressionStatement struct {
	// the expression
	Expression
}

func (es *ExpressionStatement) statementNode() {}

// NewExpressionStatement creates an ExpressionStatement node
func NewExpressionStatement(exp Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expression: exp,
	}
}

// BlockStatement represents a series of statments grouped by {}
type BlockStatement struct {
	// the { token
	Token token.Token
	// a series of statements grouped by {}
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	builder := strings.Builder{}

	for _, s := range bs.Statements {
		builder.WriteString(s.String())
	}

	return builder.String()
}

// NewBlockStatement creates a BlockStatement node
func NewBlockStatement(statements ...Statement) *BlockStatement {
	return &BlockStatement{
		Token:      token.New(token.LBRACE, "{"),
		Statements: statements,
	}
}
