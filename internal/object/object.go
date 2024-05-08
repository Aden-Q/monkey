package object

import (
	"strconv"
	"strings"

	"github.com/aden-q/monkey/internal/ast"
)

// interface compliance check
var _ Object = (*Integer)(nil)
var _ Object = (*Boolean)(nil)
var _ Object = (*Nil)(nil)
var _ Object = (*ReturnValue)(nil)
var _ Object = (*Error)(nil)
var _ Object = (*Func)(nil)

type ObjectType string

var (
	INTEGER_OBJ      = ObjectType("INTEGER")
	BOOLEAN_OBJ      = ObjectType("BOOLEAN")
	NIL_OBJ          = ObjectType("NIL")
	RETURN_VALUE_OBJ = ObjectType("RETURN_VALUE")
	ERROR_OBJ        = ObjectType("ERROR")
	FUNCTION_OBJ     = ObjectType("FUNCTION")
)

// boolean literal objects
var (
	TRUE  = NewBoolean(true)
	FALSE = NewBoolean(false)
	NIL   = NewNil()
)

type Object interface {
	Type() ObjectType
	Inspect() string
	IsTruthy() bool
}

// Integer
type Integer struct {
	Value int64
}

func NewInteger(value int64) *Integer {
	return &Integer{
		Value: value,
	}
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return strconv.FormatInt(i.Value, 10)
}

func (i *Integer) IsTruthy() bool {
	return i.Value != 0
}

type Boolean struct {
	Value bool
}

func NewBoolean(value bool) *Boolean {
	return &Boolean{
		Value: value,
	}
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

func (b *Boolean) IsTruthy() bool {
	return b.Value
}

// Nil represents the absence of any value
type Nil struct{}

func NewNil() *Nil {
	return &Nil{}
}

func (n *Nil) Type() ObjectType {
	return NIL_OBJ
}

func (n *Nil) Inspect() string {
	return "nil"
}

func (n *Nil) IsTruthy() bool {
	return false
}

// ReturnValue represents a return value of a function
type ReturnValue struct {
	Value Object
}

func NewReturnValue(value Object) *ReturnValue {
	return &ReturnValue{
		Value: value,
	}
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) IsTruthy() bool {
	return rv.Value.IsTruthy()
}

// Error represents an error
// Note: this is not necessary for the interpreter to work, we do error handling in Go's native way
type Error struct {
	Message string
}

func NewError(msg string) *Error {
	return &Error{
		Message: msg,
	}
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// FIXME: this behavior is undined, not sure an error is truthy or not
func (e *Error) IsTruthy() bool {
	return false
}

// Func represents a function object
type Func struct {
	Parameters []*ast.IdentifierExpression
	Body       *ast.BlockStatement
	// the environment for the function scope, allowing closure
	Env Environment
}

func NewFunc(params []*ast.IdentifierExpression, body *ast.BlockStatement, env Environment) *Func {
	return &Func{
		Parameters: params,
		Body:       body,
		Env:        env.Copy(),
	}
}

func (f *Func) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Func) Inspect() string {
	builder := strings.Builder{}

	paramStrings := []string{}
	for _, param := range f.Parameters {
		paramStrings = append(paramStrings, param.String())
	}

	builder.WriteString("fn")
	builder.WriteString("(")
	builder.WriteString(strings.Join(paramStrings, ", "))
	builder.WriteString(") {\n")
	builder.WriteString(f.Body.String())
	builder.WriteString("\n}")

	return builder.String()
}

// FIXME: this behavior is undined, not sure an error is truthy or not
func (f *Func) IsTruthy() bool {
	return false
}
