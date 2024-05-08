package evaluator

import (
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
	switch node.(type) {
	// parse program
	case *ast.Program:
		return e.evalStatements(node.(*ast.Program).Statements)
	// parse statement
	case *ast.ExpressionStatement:
		return e.Eval(node.(*ast.ExpressionStatement).Expression)
	// parse expression
	case *ast.IntegerExpression:
		return &object.Integer{Value: node.(*ast.IntegerExpression).Value}, nil
	case *ast.BooleanExpression:
		return getBooleanObject(node.(*ast.BooleanExpression).Value), nil
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node.(*ast.PrefixExpression))
	case *ast.InfixExpression:
		return e.evalInfixExpression(node.(*ast.InfixExpression))
	}

	return object.NIL, nil
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

// TODO: check when the ! operator can fail and return a proper error
func (e *evaluator) evalBangOperatorExpression(o object.Object) (object.Object, error) {
	switch o {
	case object.FALSE, object.NewInteger(0):
		return object.TRUE, nil
	default:
		return object.FALSE, nil
	}
}

func (e *evaluator) evalMinuxPrefixOperatorExpression(o object.Object) (object.Object, error) {
	if o.Type() != object.INTEGER_OBJ {
		return object.NIL, ErrUnexpectedObjectType
	}

	return object.NewInteger(-o.(*object.Integer).Value), nil
}

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
		return e.evalIntegerInfixExpression(ie.Operator, leftOperandObj, rightOperandObj)
	default:
		return object.NIL, ErrUnexpectedObjectType
	}
}

func (e *evaluator) evalIntegerInfixExpression(operator string, left, right object.Object) (object.Object, error) {
	leftVal, rightVal := left.(*object.Integer).Value, right.(*object.Integer).Value

	switch operator {
	case "+":
		return object.NewInteger(leftVal + rightVal), nil
	case "-":
		return object.NewInteger(leftVal - rightVal), nil
	case "*":
		return object.NewInteger(leftVal * rightVal), nil
	case "/":
		return object.NewInteger(leftVal / rightVal), nil
	default:
		return object.NIL, ErrUnexpectedOperatorType
	}
}

func getBooleanObject(input bool) object.Object {
	if input {
		return object.TRUE
	}

	return object.FALSE
}
