package object

import "strconv"

// interface compliance check
var _ Object = (*Integer)(nil)
var _ Object = (*Boolean)(nil)
var _ Object = (*Nil)(nil)

type ObjectType string

var (
	INTEGER_OBJ      = ObjectType("INTEGER")
	BOOLEAN_OBJ      = ObjectType("BOOLEAN")
	NIL_OBJ          = ObjectType("NIL")
	RETURN_VALUE_OBJ = ObjectType("RETURN_VALUE")
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

// ReturnValue represents the return value of a function
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
