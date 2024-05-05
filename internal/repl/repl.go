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

	fmt.Print(MONKEY_FACE)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", userName)

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

		// TODO: PrettyPrint
		fmt.Println(program.String())
	}
}

func printParserErrors(out io.WriteCloser, errs []error) {
	fmt.Println("parser errors:")

	for _, err := range errs {
		fmt.Println("\t" + err.Error())
	}
}
