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
	l := lexer.New()

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l.Read(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
