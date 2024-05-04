package parser_test

import (
	"github.com/Aden-Q/monkey/internal/ast"
	"github.com/Aden-Q/monkey/internal/lexer"
	"github.com/Aden-Q/monkey/internal/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var (
		text    string
		l       lexer.Lexer
		p       parser.Parser
		program *ast.Program
		errs    []error
	)

	BeforeEach(func() {
		l = lexer.New()
		p = parser.New(l)
	})

	Describe("ParseProgram", func() {
		Context("parse expressions", func() {
			It("identifier expressions", func() {
				text = `
				foo;
				bar;
				cs;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewIdentifier("foo")),
						ast.NewExpressionStatement(ast.NewIdentifier("bar")),
						ast.NewExpressionStatement(ast.NewIdentifier("cs")),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("integer expressions", func() {
				text = `
				5;
				10;
				838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewInteger("5", 5)),
						ast.NewExpressionStatement(ast.NewInteger("10", 10)),
						ast.NewExpressionStatement(ast.NewInteger("838383", 838383)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})
		})

		Context("parse let statements", func() {
			It("correct program", func() {
				text = `
				let x = 5;
				let y = 10;
				let foobar = 838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewLetStatement(
							ast.NewIdentifier("x"),
							ast.NewInteger("5", 5),
						),
						ast.NewLetStatement(
							ast.NewIdentifier("y"),
							ast.NewInteger("10", 10),
						),
						ast.NewLetStatement(
							ast.NewIdentifier("foobar"),
							ast.NewInteger("838383", 838383),
						),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("missing identifiers", func() {
				text = `
				let = 5;
				let = 10;
				let = 838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{},
				}
				expectedErrors := []error{
					parser.ErrUnexpectedTokenType,
					parser.ErrUnexpectedTokenType,
					parser.ErrUnexpectedTokenType,
				}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("missing assign tokens", func() {
				text = `
				let x 5;
				let y 10;
				let foo 838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{},
				}
				expectedErrors := []error{
					parser.ErrUnexpectedTokenType,
					parser.ErrUnexpectedTokenType,
					parser.ErrUnexpectedTokenType,
				}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})
		})

		Context("parse return statements", func() {
			It("can parse the program when there is no error", func() {
				text = `
				return 5;
				return 10;
				return 838383;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewReturnStatement(ast.NewInteger("5", 5)),
						ast.NewReturnStatement(ast.NewInteger("10", 10)),
						ast.NewReturnStatement(ast.NewInteger("838383", 838383)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})
		})
	})
})
