package object

import (
	"errors"
)

var (
	ErrWrongNumberArguments    = errors.New("wrong number of argument(s)")
	ErrUnsupportedArgumentType = errors.New("unsupported argument type")
)
