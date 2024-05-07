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

// Evaluate recursively evaluate an AST node
func (e *evaluator) Eval(node ast.Node) (object.Object, error) {
	switch node.(type) {
	case *ast.Program:
		return e.evalStatements(node.(*ast.Program).Statements)
	case *ast.ExpressionStatement:
		return e.Eval(node.(*ast.ExpressionStatement).Expression)
	case *ast.IntegerExpression:
		return &object.Integer{Value: node.(*ast.IntegerExpression).Value}, nil
	case *ast.BooleanExpression:
		return &object.Boolean{Value: node.(*ast.BooleanExpression).Value}, nil
	}

	return nil, nil
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
