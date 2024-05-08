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
}

type evaluator struct {
}

func New() Evaluator {
	return &evaluator{}
}

// Eval evaluate an AST node recursively
func (e *evaluator) Eval(node ast.Node) (object.Object, error) {
	if node == nil {
		return object.NIL, ErrEmptyNodeInput
	}

	switch node := node.(type) {
	// evaluate the program
	case *ast.Program:
		return e.evalStatements(node.Statements)
	// evaluate statements
	case *ast.ExpressionStatement:
		return e.Eval(node.Expression)
	case *ast.BlockStatement:
		return e.evalStatements(node.Statements)
	// evaluate expressions
	case *ast.IntegerExpression:
		return object.NewInteger(node.Value), nil
	case *ast.BooleanExpression:
		return booleanConv(node.Value), nil
	case *ast.IfExpression:
		return e.evalIfExpression(node)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node)
	case *ast.InfixExpression:
		return e.evalInfixExpression(node)
	}

	// no match, unexpected path
	return object.NIL, ErrUnexpectedNodeType
}

func (e *evaluator) evalStatements(stmts []ast.Statement) (object.Object, error) {
	var result object.Object
	var err error

	// iteratively evaluate each statement and return the result from the last one
	for _, stmt := range stmts {
		result, err = e.Eval(stmt)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *evaluator) evalIfExpression(ie *ast.IfExpression) (object.Object, error) {
	condition, err := e.Eval(ie.Condition)
	if err != nil {
		return object.NIL, nil
	}

	if condition.IsTruthy() {
		return e.Eval(ie.Consequence)
	}

	if ie.Alternative != nil {
		return e.Eval(ie.Alternative)
	}

	return object.NIL, nil
}

func (e *evaluator) evalPrefixExpression(pe *ast.PrefixExpression) (object.Object, error) {
	operandObj, err := e.Eval(pe.Operand)
	if err != nil {
		return nil, nil
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
		return nil, err
	}

	rightOperandObj, err := e.Eval(ie.RightOperand)
	if err != nil {
		return nil, err
	}

	switch {
	case leftOperandObj.Type() == object.INTEGER_OBJ && rightOperandObj.Type() == object.INTEGER_OBJ:
		return e.evalIntegerInfixExpression(ie.Operator, leftOperandObj.(*object.Integer), rightOperandObj.(*object.Integer))
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

// evalIntegerInfixExpression evaluates an infix expression involving two integer operators
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

// booleanConv converts a boolean literal to a boolean object in the object system
func booleanConv(input bool) object.Object {
	if input {
		return object.TRUE
	}

	return object.FALSE
}
