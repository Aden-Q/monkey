package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Aden-Q/monkey/internal/lexer"
	"github.com/Aden-Q/monkey/internal/token"
)

const PROMPT = ">>> "

type REPL interface {
	Start(in io.ReadCloser, out io.WriteCloser)
}

type Config struct {
}

type repl struct {
}

func New(config Config) REPL {
	return &repl{}
}

func (r *repl) Start(in io.ReadCloser, out io.WriteCloser) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok, ok := l.NextToken(); ok && tok.Type != token.EOF; tok, ok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
