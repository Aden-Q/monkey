package evaluator

import (
	"errors"
)

var (
	ErrEmptyNodeInput         = errors.New("empty node input")
	ErrUnexpectedNodeType     = errors.New("unexpected node type")
	ErrUnexpectedObjectType   = errors.New("unexpected object type")
	ErrUnexpectedOperatorType = errors.New("unexpected operator type")
	ErrIdentifierNotFound     = errors.New("identifier not found")
	ErrNotAFunction           = errors.New("not a function")
	ErrIndexOutOfRange        = errors.New("index out of range")
	ErrUnhashableType         = errors.New("unhashable type")
	ErrKeyNotFound            = errors.New("key not found")
)
