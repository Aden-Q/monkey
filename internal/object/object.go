package object

import (
	"hash/fnv"
	"strconv"
	"strings"

	"github.com/aden-q/monkey/internal/ast"
)

// interface compliance check
var _ Object = (*Integer)(nil)
var _ Object = (*Boolean)(nil)
var _ Object = (*String)(nil)
var _ Object = (*Array)(nil)
var _ Object = (*Hash)(nil)
var _ Object = (*Nil)(nil)
var _ Object = (*ReturnValue)(nil)
var _ Object = (*Error)(nil)
var _ Object = (*Func)(nil)
var _ Object = (BuiltinFunc)(nil)

type ObjectType string

var (
	INTEGER_OBJ      = ObjectType("INTEGER")
	BOOLEAN_OBJ      = ObjectType("BOOLEAN")
	STRING_OBJ       = ObjectType("STRING")
	ARRAY_OBJ        = ObjectType("ARRAY")
	HASH_OBJ         = ObjectType("HASH")
	NIL_OBJ          = ObjectType("NIL")
	RETURN_VALUE_OBJ = ObjectType("RETURN_VALUE")
	ERROR_OBJ        = ObjectType("ERROR")
	FUNCTION_OBJ     = ObjectType("FUNCTION")
	BUILTINFUNC_OBJ  = ObjectType("BUILTINFUNC")
)

// boolean literal objects
var (
	TRUE  = NewBoolean(true)
	FALSE = NewBoolean(false)
	NIL   = NewNil()
)

var BuiltinFuncs = map[string]BuiltinFunc{
	"len": func(args ...Object) (Object, error) {
		if len(args) != 1 {
			return NIL, ErrWrongNumberArguments
		}

		switch arg := args[0].(type) {
		case *String:
			return NewInteger(int64(len(arg.Value))), nil
		case *Array:
			return NewInteger(int64(len(arg.Elements))), nil
		default:
			return NIL, ErrUnsupportedArgumentType
		}
	},
}

type Object interface {
	Type() ObjectType
	Inspect() string
	IsTruthy() bool
}

type HashKey struct {
	Type   ObjectType
	Object Object
	Value  uint64
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

func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type:   i.Type(),
		Object: i,
		Value:  uint64(i.Value),
	}
}

// the boolean object
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

func (b *Boolean) HashKey() HashKey {
	var val uint64 = 0
	if b.Value {
		val = 1
	}

	return HashKey{
		Type:   b.Type(),
		Object: b,
		Value:  val,
	}
}

// the string object
type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{
		Value: value,
	}
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) IsTruthy() bool {
	return len(s.Value) > 0
}

func (s *String) HashKey() HashKey {
	hash := fnv.New64a()
	hash.Write([]byte(s.Value))

	return HashKey{
		Type:   s.Type(),
		Object: s,
		Value:  hash.Sum64(),
	}
}

// the array object
type Array struct {
	Elements []Object
}

func NewArray(elements ...Object) *Array {
	return &Array{
		Elements: elements,
	}
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (a *Array) Inspect() string {
	builder := strings.Builder{}

	elements := []string{}
	for _, element := range a.Elements {
		elements = append(elements, element.Inspect())
	}

	builder.WriteString("[")
	builder.WriteString(strings.Join(elements, ", "))
	builder.WriteString("]")

	return builder.String()
}

func (a *Array) IsTruthy() bool {
	return len(a.Elements) > 0
}

// the hash object
type Hash struct {
	Items map[HashKey]Object
}

func NewHash(items map[HashKey]Object) *Hash {
	return &Hash{
		Items: items,
	}
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {
	builder := strings.Builder{}

	items := []string{}
	for key, value := range h.Items {
		items = append(items, key.Object.Inspect()+": "+value.Inspect())
	}

	builder.WriteString("{")
	builder.WriteString(strings.Join(items, ", "))
	builder.WriteString("}")

	return builder.String()
}

func (h *Hash) IsTruthy() bool {
	return len(h.Items) > 0
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
}

func NewFunc(params []*ast.IdentifierExpression, body *ast.BlockStatement, env Environment) *Func {
	return &Func{
		Parameters: params,
		Body:       body,
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

// FIXME: this behavior is undined, not sure whether a function is truthy or not
func (f *Func) IsTruthy() bool {
	return false
}

// BuiltinFunc represents a builtin function object
type BuiltinFunc func(args ...Object) (Object, error)

func (b BuiltinFunc) Type() ObjectType {
	return FUNCTION_OBJ
}

func (b BuiltinFunc) Inspect() string {
	return "builtin function"
}

// FIXME: this behavior is undined, not sure whether a builtin function is truthy or not
func (b BuiltinFunc) IsTruthy() bool {
	return false
}
