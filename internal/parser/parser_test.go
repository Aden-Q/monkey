package parser_test

import (
	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
	"github.com/Aden-Q/monkey/internal/parser"
	"github.com/Aden-Q/monkey/internal/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var (
		text    string
		l       lexer.Lexer
		p       parser.Parser
		program *ast.Program
	)

	BeforeEach(func() {
		l = lexer.New()
		p = parser.New(l)
	})

	Describe("Parser", func() {
		Context("ParseProgram", func() {
			It("can parse the program", func() {
				text = `
				let x = 5;
				let y = 10;
				let foobar = 838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewLetStatement(&ast.Identifier{
							Token: token.New(token.IDENT, "x"),
						}, nil),
						ast.NewLetStatement(&ast.Identifier{
							Token: token.New(token.IDENT, "y"),
						}, nil),
						ast.NewLetStatement(&ast.Identifier{
							Token: token.New(token.IDENT, "foobar"),
						}, nil),
					},
				}

				program = p.ParseProgram(text)
				Expect(program).ToNot(BeNil())
				// expect to have 3 let statments
				Expect(len(program.Statements)).To(Equal(3))
				// expect deep equal
				Expect(program).To(Equal(expectedProgram))
			})
		})
	})
})
