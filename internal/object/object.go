package object

import "strconv"

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NIL_OBJ     = "NIL"
)

// interface compliance check
var _ Object = (*Integer)(nil)
var _ Object = (*Boolean)(nil)
var _ Object = (*Nil)(nil)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return strconv.FormatInt(i.Value, 10)
}

func NewInteger(value int64) *Integer {
	return &Integer{
		Value: value,
	}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

func NewBoolean(value bool) *Boolean {
	return &Boolean{
		Value: value,
	}
}

// Nil represents the absence of any value
type Nil struct {
}

func (n *Nil) Type() ObjectType {
	return NIL_OBJ
}

func (n *Nil) Inspect() string {
	return "nil"
}

func NewNil(value bool) *Boolean {
	return &Boolean{
		Value: value,
	}
}
