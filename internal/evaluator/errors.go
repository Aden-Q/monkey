package evaluator

import (
	"errors"
)

var (
	ErrUnexpectedObjectType   = errors.New("unexpected object type")
	ErrUnexpectedOperatorType = errors.New("unexpected operator type")
)
