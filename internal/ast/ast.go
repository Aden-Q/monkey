package ast

import (
	"strings"

	"github.com/aden-q/monkey/internal/token"
)

// interface compliance check
var _ Node = (*Program)(nil)
var _ Expression = (*IdentifierExpression)(nil)
var _ Expression = (*IntegerExpression)(nil)
var _ Expression = (*BooleanExpression)(nil)
var _ Expression = (*StringExpression)(nil)
var _ Expression = (*ArrayExpression)(nil)
var _ Expression = (*IfExpression)(nil)
var _ Expression = (*FuncExpression)(nil)
var _ Expression = (*CallExpression)(nil)
var _ Expression = (*PrefixExpression)(nil)
var _ Expression = (*InfixExpression)(nil)
var _ Statement = (*LetStatement)(nil)
var _ Statement = (*ReturnStatement)(nil)
var _ Statement = (*ExpressionStatement)(nil)
var _ Statement = (*BlockStatement)(nil)

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
	if statements == nil {
		return &Program{
			Statements: []Statement{},
		}
	}

	return &Program{
		Statements: statements,
	}
}

// -------------- Expressions -------------------

// IdentifierExpression implements the Expression interface
type IdentifierExpression struct {
	// the identifier token
	Token token.Token
	Value string
}

func (ie *IdentifierExpression) expressionNode() {}

func (ie *IdentifierExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IdentifierExpression) String() string {
	return ie.Value
}

// NewIdentifierExpression creates an Identifier node
func NewIdentifierExpression(literal string) *IdentifierExpression {
	return &IdentifierExpression{
		Token: token.New(token.IDENT, literal),
		Value: literal,
	}
}

// IntegerExpression implements the Expression interface
type IntegerExpression struct {
	// the integer token
	Token token.Token
	Value int64
}

func (ie *IntegerExpression) expressionNode() {}

func (ie *IntegerExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IntegerExpression) String() string {
	return ie.Token.Literal
}

// NewIntegerExpression creates an Integer node
func NewIntegerExpression(literal string, value int64) *IntegerExpression {
	return &IntegerExpression{
		Token: token.New(token.INT, literal),
		Value: value,
	}
}

// BooleanExpression implements the Expression interface
type BooleanExpression struct {
	// the boolean token
	Token token.Token
	Value bool
}

func (be *BooleanExpression) expressionNode() {}

func (be *BooleanExpression) TokenLiteral() string {
	return be.Token.Literal
}

func (be *BooleanExpression) String() string {
	return be.Token.Literal
}

// NewBooleanExpression creates a Boolean node
func NewBooleanExpression(value bool) *BooleanExpression {
	var tok token.Token
	if value {
		tok = token.New(token.TRUE, "true")
	} else {
		tok = token.New(token.FALSE, "false")
	}

	return &BooleanExpression{
		Token: tok,
		Value: value,
	}
}

// StringExpression implements the Expression interface
type StringExpression struct {
	// the string token
	Token token.Token
	Value string
}

func (se *StringExpression) expressionNode() {}

func (se *StringExpression) TokenLiteral() string {
	return se.Token.Literal
}

func (se *StringExpression) String() string {
	return se.Token.Literal
}

// NewStringExpression creates a String node
func NewStringExpression(literal string) *StringExpression {
	return &StringExpression{
		Token: token.New(token.STRING, literal),
		Value: literal,
	}
}

// ArrayExpression implements the Expression interface
type ArrayExpression struct {
	// the [ token
	Token    token.Token
	Elements []Expression
}

func (ae *ArrayExpression) expressionNode() {}

func (ae *ArrayExpression) TokenLiteral() string {
	return ae.Token.Literal
}

func (ae *ArrayExpression) String() string {
	builder := strings.Builder{}

	elements := []string{}
	for _, el := range ae.Elements {
		elements = append(elements, el.String())
	}

	builder.WriteString("[")
	builder.WriteString(strings.Join(elements, ", "))
	builder.WriteString("]")

	return builder.String()
}

// NewArrayExpression creates a String node
func NewArrayExpression(exps ...Expression) *ArrayExpression {
	return &ArrayExpression{
		Token:    token.New(token.LBRACKET, "["),
		Elements: exps,
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

// FuncExpression implements the Expression interface
type FuncExpression struct {
	// the fn token
	Token token.Token
	// function parameters
	Parameters []*IdentifierExpression
	// function body
	Body *BlockStatement
}

func (fe *FuncExpression) expressionNode() {}

func (fe *FuncExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *FuncExpression) String() string {
	builder := strings.Builder{}

	paramStrings := []string{}
	for _, param := range fe.Parameters {
		paramStrings = append(paramStrings, param.String())
	}

	builder.WriteString("fn")
	builder.WriteString("(")
	builder.WriteString(strings.Join(paramStrings, ", "))
	builder.WriteString(") {\n")
	builder.WriteString(fe.Body.String())
	builder.WriteString("\n")

	return builder.String()
}

// NewFuncExpression creates a FuncExpression node
func NewFuncExpression(params []*IdentifierExpression, body *BlockStatement) *FuncExpression {
	return &FuncExpression{
		Token:      token.New(token.FUNC, "fn"),
		Parameters: params,
		Body:       body,
	}
}

// CallExpression implements the Expression interface
type CallExpression struct {
	// the first token (fn or the identifier)
	Token token.Token
	// function literal or identifier bound to function
	Func Expression
	// function call arguments
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	builder := strings.Builder{}

	argStrings := []string{}
	for _, arg := range ce.Arguments {
		argStrings = append(argStrings, arg.String())
	}

	builder.WriteString(ce.Func.String())
	builder.WriteString("(")
	builder.WriteString(strings.Join(argStrings, ", "))
	builder.WriteString(")")

	return builder.String()
}

// NewCallExpression creates a CallExpression node
func NewCallExpression(fn Expression, args []Expression) *CallExpression {
	if fn.TokenLiteral() == "fn" {
		return &CallExpression{
			Token:     token.New(token.FUNC, "fn"),
			Func:      fn,
			Arguments: args,
		}
	}

	return &CallExpression{
		Token:     token.New(token.IDENT, fn.TokenLiteral()),
		Func:      fn,
		Arguments: args,
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
	Identifier *IdentifierExpression
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
func NewLetStatement(identifier *IdentifierExpression, value Expression) *LetStatement {
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
	if statements == nil {
		return &BlockStatement{
			Token:      token.New(token.LBRACE, "{"),
			Statements: []Statement{},
		}
	}

	return &BlockStatement{
		Token:      token.New(token.LBRACE, "{"),
		Statements: statements,
	}
}
