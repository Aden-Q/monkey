package parser

import (
	"errors"
)

var (
	ErrUnexpectedTokenType   = errors.New("unexpected token type")
	ErrPrefixParseFnNotFound = errors.New("prefix parse function not found")
	ErrInfixParseFnNotFound  = errors.New("infix parse function not found")
)
