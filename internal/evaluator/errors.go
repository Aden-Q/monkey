package evaluator

import (
	"errors"
)

var (
	ErrEmptyNodeInput         = errors.New("empty node input")
	ErrUnexpectedNodeType     = errors.New("unexpected node type")
	ErrUnexpectedObjectType   = errors.New("unexpected object type")
	ErrUnexpectedOperatorType = errors.New("unexpected operator type")
)
