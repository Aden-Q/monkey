package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aden-q/monkey/internal/lexer"
	"github.com/aden-q/monkey/internal/parser"
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
	p := parser.New(l)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		program, errs := p.ParseProgram(line)
		if len(errs) != 0 {
			printParserErrors(out, errs)
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.WriteCloser, errs []error) {
	for _, err := range errs {
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
