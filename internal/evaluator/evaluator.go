package evaluator

import (
	"reflect"

	"github.com/aden-q/monkey/internal/ast"
	"github.com/aden-q/monkey/internal/object"
)

// interface compliance check
var _ Evaluator = (*evaluator)(nil)

// An interpreter/evaluator interface
type Evaluator interface {
	Eval(node ast.Node) (object.Object, error)
	extendFunctionEnv(fn *object.Func, args []object.Object)
}

type evaluator struct {
	env object.Environment
}

func New(env object.Environment) Evaluator {
	return &evaluator{
		env: env,
	}
}

// Eval evaluate an AST node recursively
func (e *evaluator) Eval(node ast.Node) (object.Object, error) {
	if node == nil {
		return object.NIL, ErrEmptyNodeInput
	}

	// explicitly polymorphic switch statement
	switch node := node.(type) {
	// evaluate the program
	case *ast.Program:
		return e.evalStatements(node.Statements)
	// evaluate statements
	case *ast.ExpressionStatement:
		return e.Eval(node.Expression)
	case *ast.BlockStatement:
		return e.evalStatements(node.Statements)
	case *ast.LetStatement:
		return e.evalLetStatement(node)
	case *ast.ReturnStatement:
		return e.evalReturnStatement(node)
	// evaluate expressions
	case *ast.IdentifierExpression:
		return e.evalIdentifierExpression(node)
	case *ast.IntegerExpression:
		return object.NewInteger(node.Value), nil
	case *ast.BooleanExpression:
		return booleanConv(node.Value), nil
	case *ast.StringExpression:
		return object.NewString(node.Value), nil
	case *ast.ArrayExpression:
		return e.evalArrayExpression(node)
	case *ast.HashExpression:
		return e.evalHashExpression(node)
	case *ast.IndexExpression:
		return e.evalIndexExpression(node)
	case *ast.IfExpression:
		return e.evalIfExpression(node)
	case *ast.FuncExpression:
		return e.evalFuncExpression(node)
	case *ast.CallExpression:
		return e.evalCallExpression(node)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node)
	case *ast.InfixExpression:
		return e.evalInfixExpression(node)
	}

	// no match, unexpected path
	return object.NIL, ErrUnexpectedNodeType
}

func (e *evaluator) evalStatements(stmts []ast.Statement) (object.Object, error) {
	var result object.Object = object.NIL
	var err error

	// iteratively evaluate each statement and return the result from the last one
	for _, stmt := range stmts {
		result, err = e.Eval(stmt)
		if err != nil {
			return object.NIL, err
		}

		// short-circuit return statement
		if result.Type() == object.RETURN_VALUE_OBJ {
			return result, nil
		}
	}

	return result, nil
}

func (e *evaluator) evalLetStatement(stmt *ast.LetStatement) (object.Object, error) {
	val, err := e.Eval(stmt.Value)
	if err != nil {
		return object.NIL, err
	}

	// bind the evaluated value to the environment
	e.env.Set(stmt.Identifier.Value, val)

	return object.NIL, nil
}

func (e *evaluator) evalReturnStatement(stmt *ast.ReturnStatement) (object.Object, error) {
	val, err := e.Eval(stmt.Value)
	if err != nil {
		return object.NIL, err
	}

	return object.NewReturnValue(val), nil
}

func (e *evaluator) evalIdentifierExpression(ie *ast.IdentifierExpression) (object.Object, error) {
	if val, ok := e.env.Get(ie.Value); ok {
		return val, nil
	}

	if builtinFunc, ok := object.BuiltinFuncs[ie.Value]; ok {
		return builtinFunc, nil
	}

	return object.NIL, ErrIdentifierNotFound
}

func (e *evaluator) evalArrayExpression(ae *ast.ArrayExpression) (object.Object, error) {
	elements := make([]object.Object, 0, len(ae.Elements))

	for _, exp := range ae.Elements {
		val, err := e.Eval(exp)
		if err != nil {
			return object.NIL, err
		}

		elements = append(elements, val)
	}

	return object.NewArray(elements...), nil
}

func (e *evaluator) evalHashExpression(he *ast.HashExpression) (object.Object, error) {
	items := make(map[object.HashKey]object.Object)

	for keyNode, valueNode := range he.Items {
		key, err := e.Eval(keyNode)
		if err != nil {
			return object.NIL, err
		}

		value, err := e.Eval(valueNode)
		if err != nil {
			return object.NIL, err
		}

		// a key must be hashable in order to be used as a key in a hash object
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return object.NIL, ErrUnhashableType
		}

		items[hashKey.HashKey()] = value
	}

	return object.NewHash(items), nil
}

func (e *evaluator) evalIndexExpression(ae *ast.IndexExpression) (object.Object, error) {
	left, err := e.Eval(ae.Left)
	if err != nil {
		return object.NIL, err
	}

	index, err := e.Eval(ae.Index)
	if err != nil {
		return object.NIL, err
	}

	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		array := left.(*object.Array)
		idx := index.(*object.Integer).Value
		maxIdx := int64(len(array.Elements) - 1)

		if idx < 0 || idx > maxIdx {
			return object.NIL, ErrIndexOutOfRange
		}

		return array.Elements[idx], nil
	case left.Type() == object.HASH_OBJ:
		hash := left.(*object.Hash)

		// a key must be hashable in order to be used as a key in a hash object
		hashKey, ok := index.(object.Hashable)
		if !ok {
			return object.NIL, ErrUnhashableType
		}

		if val, ok := hash.Items[hashKey.HashKey()]; ok {
			return val, nil
		}

		return object.NIL, ErrKeyNotFound
	default:
		return object.NIL, ErrUnexpectedObjectType
	}
}

func (e *evaluator) evalIfExpression(ie *ast.IfExpression) (object.Object, error) {
	condition, err := e.Eval(ie.Condition)
	if err != nil {
		return object.NIL, err
	}

	if condition.IsTruthy() {
		return e.Eval(ie.Consequence)
	}

	if ie.Alternative != nil {
		return e.Eval(ie.Alternative)
	}

	return object.NIL, nil
}

func (e *evaluator) evalFuncExpression(fe *ast.FuncExpression) (object.Object, error) {
	return object.NewFunc(fe.Parameters, fe.Body, e.env), nil
}

func (e *evaluator) evalCallExpression(ce *ast.CallExpression) (object.Object, error) {
	function, err := e.Eval(ce.Func)
	if err != nil {
		return object.NIL, err
	}

	args := make([]object.Object, 0, len(ce.Arguments))

	for _, exp := range ce.Arguments {
		res, err := e.Eval(exp)
		if err != nil {
			return object.NIL, err
		}

		args = append(args, res)
	}

	// call the function with the given arguments
	return applyFunc(function, args, e.env)
}

func applyFunc(fn object.Object, args []object.Object, env object.Environment) (object.Object, error) {
	switch fn := fn.(type) {
	case *object.Func:
		// create a new environment for the function scope
		funcEvaluator := New(object.NewClosureEnvironment(env))
		// extend the closure environment with arguments passed to the function
		funcEvaluator.extendFunctionEnv(fn, args)

		val, err := funcEvaluator.Eval(fn.Body)
		if err != nil {
			return object.NIL, err
		}

		if returnVal, ok := val.(*object.ReturnValue); ok {
			return returnVal.Value, nil
		}

		return val, nil
	case object.BuiltinFunc:
		return fn(args...)
	default:
		return object.NIL, ErrNotAFunction
	}
}

func (e *evaluator) extendFunctionEnv(fn *object.Func, args []object.Object) {
	for i, param := range fn.Parameters {
		e.env.Set(param.Value, args[i])
	}
}

func (e *evaluator) evalPrefixExpression(pe *ast.PrefixExpression) (object.Object, error) {
	operandObj, err := e.Eval(pe.Operand)
	if err != nil {
		return object.NIL, err
	}

	switch pe.Operator {
	case "!":
		return e.evalBangOperatorExpression(operandObj)
	case "-":
		return e.evalMinuxPrefixOperatorExpression(operandObj)
	default:
		return object.NIL, nil
	}
}

// evalBangOperatorExpression evaluates a prefix expression with a '!' token as the prefix
func (e *evaluator) evalBangOperatorExpression(o object.Object) (object.Object, error) {
	switch o {
	case object.FALSE, object.NewInteger(0):
		return object.TRUE, nil
	default:
		return object.FALSE, nil
	}
}

// evalMinuxPrefixOperatorExpression evaluates a prefix expression with a '-' token as the prefix
func (e *evaluator) evalMinuxPrefixOperatorExpression(o object.Object) (object.Object, error) {
	if o.Type() != object.INTEGER_OBJ {
		return object.NIL, ErrUnexpectedObjectType
	}

	return object.NewInteger(-o.(*object.Integer).Value), nil
}

// evalInfixExpression evaluates an infix expression
func (e *evaluator) evalInfixExpression(ie *ast.InfixExpression) (object.Object, error) {
	leftOperandObj, err := e.Eval(ie.LeftOperand)
	if err != nil {
		return object.NIL, err
	}

	rightOperandObj, err := e.Eval(ie.RightOperand)
	if err != nil {
		return object.NIL, err
	}

	switch {
	case leftOperandObj.Type() == object.INTEGER_OBJ && rightOperandObj.Type() == object.INTEGER_OBJ:
		return e.evalIntegerInfixExpression(ie.Operator, leftOperandObj.(*object.Integer), rightOperandObj.(*object.Integer))
	case leftOperandObj.Type() == object.STRING_OBJ && rightOperandObj.Type() == object.STRING_OBJ:
		return e.evalStringInfixExpression(ie.Operator, leftOperandObj.(*object.String), rightOperandObj.(*object.String))
	// equality test
	case ie.Operator == "==":
		return booleanConv(reflect.DeepEqual(leftOperandObj, rightOperandObj)), nil
	case ie.Operator == "!=":
		return booleanConv(!reflect.DeepEqual(leftOperandObj, rightOperandObj)), nil
	default:
		// TODO: check infix expressions involving boolean operands and operators that result in boolean values
		return object.NIL, ErrUnexpectedObjectType
	}
}

// evalIntegerInfixExpression evaluates an infix expression involving two integer operands and a single operator
func (e *evaluator) evalIntegerInfixExpression(operator string, left, right *object.Integer) (object.Object, error) {
	leftVal, rightVal := left.Value, right.Value

	switch operator {
	case "+":
		return object.NewInteger(leftVal + rightVal), nil
	case "-":
		return object.NewInteger(leftVal - rightVal), nil
	case "*":
		return object.NewInteger(leftVal * rightVal), nil
	case "/":
		return object.NewInteger(leftVal / rightVal), nil
	case "<":
		return booleanConv(leftVal < rightVal), nil
	case "<=":
		return booleanConv(leftVal <= rightVal), nil
	case ">":
		return booleanConv(leftVal > rightVal), nil
	case ">=":
		return booleanConv(leftVal >= rightVal), nil
	case "==":
		return booleanConv(leftVal == rightVal), nil
	case "!=":
		return booleanConv(leftVal != rightVal), nil
	default:
		return object.NIL, ErrUnexpectedOperatorType
	}
}

// evalStringInfixExpression evaluates an infix expression involving two string operands and a single operator
func (e *evaluator) evalStringInfixExpression(operator string, left, right *object.String) (object.Object, error) {
	leftVal, rightVal := left.Value, right.Value

	switch operator {
	case "+":
		return object.NewString(leftVal + rightVal), nil
	}

	return object.NIL, ErrUnexpectedOperatorType
}

// booleanConv converts a boolean literal to a boolean object in the object system
func booleanConv(input bool) object.Object {
	if input {
		return object.TRUE
	}

	return object.FALSE
}
