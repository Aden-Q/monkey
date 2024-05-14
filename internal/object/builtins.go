package object

import (
	"fmt"
	"strings"
)

// all built-in functions
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
	"print": func(args ...Object) (Object, error) {
		strToPrint := []string{}
		for _, arg := range args {
			strToPrint = append(strToPrint, arg.Inspect())
		}

		fmt.Println(strings.Join(strToPrint, " "))

		return NIL, nil
	},
}
