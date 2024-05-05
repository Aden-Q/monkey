package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aden-q/monkey/internal/lexer"
	"github.com/aden-q/monkey/internal/parser"
)

const PROMPT = ">>> "

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

type REPL interface {
	Start(in io.ReadCloser, out io.WriteCloser, userName string)
}

type Config struct {
}

type repl struct {
}

func New(config Config) REPL {
	return &repl{}
}

func (r *repl) Start(in io.ReadCloser, out io.WriteCloser, userName string) {
	scanner := bufio.NewScanner(in)
	l := lexer.New()
	p := parser.New(l)

	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, fmt.Sprintf("Hello %s! This is the Monkey programming language!\n", userName))

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
	io.WriteString(out, "parser errors:\n")

	for _, err := range errs {
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
