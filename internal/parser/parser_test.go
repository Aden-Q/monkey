package parser_test

import (
	"github.com/aden-q/monkey/internal/ast"
	"github.com/aden-q/monkey/internal/lexer"
	"github.com/aden-q/monkey/internal/parser"
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
			It("simple identifier expressions", func() {
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

			It("simple integer expressions", func() {
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

			It("simple boolean expressions", func() {
				text = `
				true;
				false;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewBoolean(true)),
						ast.NewExpressionStatement(ast.NewBoolean(false)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("simple if expression", func() {
				text = `
				if (x < y) { x; };
				if (x < y) { x; } else { y; };
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewIfExpression(
							ast.NewInfixExpression("<", ast.NewIdentifier("x"), ast.NewIdentifier("y")),
							ast.NewBlockStatement(ast.NewExpressionStatement(ast.NewIdentifier("x"))),
							nil,
						)),
						ast.NewExpressionStatement(ast.NewIfExpression(
							ast.NewInfixExpression("<", ast.NewIdentifier("x"), ast.NewIdentifier("y")),
							ast.NewBlockStatement(ast.NewExpressionStatement(ast.NewIdentifier("x"))),
							ast.NewBlockStatement(ast.NewExpressionStatement(ast.NewIdentifier("y"))),
						)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("simple prefix expressions", func() {
				text = `
				-5;
				!10;
				-foo;
				!bar;
				!true;
				!false;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewPrefixExpression("-", ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewInteger("10", 10))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("-", ast.NewIdentifier("foo"))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewIdentifier("bar"))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewBoolean(true))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewBoolean(false))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("simple infix expressions", func() {
				text = `
				5 + 5;
				5 - 5;
				5 * 5;
				5 / 5;
				6 > 5;
				6 >= 5;
				5 < 6;
				5 <= 6;
				5 == 5;
				5 != 6;
				foo + bar;
				foo + 5;
				bar + 5;
				-a * b;
				true == true;
				true != false;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewInteger("5", 5), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("-", ast.NewInteger("5", 5), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("*", ast.NewInteger("5", 5), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("/", ast.NewInteger("5", 5), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression(">", ast.NewInteger("6", 6), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression(">=", ast.NewInteger("6", 6), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("<", ast.NewInteger("5", 5), ast.NewInteger("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("<=", ast.NewInteger("5", 5), ast.NewInteger("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("==", ast.NewInteger("5", 5), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("!=", ast.NewInteger("5", 5), ast.NewInteger("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifier("foo"), ast.NewIdentifier("bar"))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifier("foo"), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifier("bar"), ast.NewInteger("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("*", ast.NewPrefixExpression("-", ast.NewIdentifier("a")), ast.NewIdentifier("b"))),
						ast.NewExpressionStatement(ast.NewInfixExpression("==", ast.NewBoolean(true), ast.NewBoolean(true))),
						ast.NewExpressionStatement(ast.NewInfixExpression("!=", ast.NewBoolean(true), ast.NewBoolean(false))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("complex infix expression string match", func() {
				texts := []string{
					`-a * b;`,
					`!-a;`,
					`!a + b;`,
					`a + b + c;`,
					`a + b * c;`,
					`a + b * c + d / e - f;`,
					`3 + 4 * 5 == 3 * 1 + 4 * 5;`,
					`3 > 5 == false;`,
					`1 + (2 + 3) + 4;`,
					`(5 + 5) * 2;`,
					`-(5 + 5);`,
					`!(true == false);`,
				}
				expectedStrings := []string{
					`((-a) * b)`,
					`(!(-a))`,
					`((!a) + b)`,
					`((a + b) + c)`,
					`(a + (b * c))`,
					`(((a + (b * c)) + (d / e)) - f)`,
					`((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))`,
					`((3 > 5) == false)`,
					`((1 + (2 + 3)) + 4)`,
					`((5 + 5) * 2)`,
					`(-(5 + 5))`,
					`(!(true == false))`,
				}
				expectedErrors := []error{}

				for idx := range texts {
					program, errs = p.ParseProgram(texts[idx])
					Expect(program.String()).To(Equal(expectedStrings[idx]))
				}

				Expect(errs).To(Equal(expectedErrors))
			})
		})

		Context("parse let statements", func() {
			It("correct program", func() {
				text = `
				let x = 5;
				let y = 10;
				let foo = 838383;
				let foo = true;
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
							ast.NewIdentifier("foo"),
							ast.NewInteger("838383", 838383),
						),
						ast.NewLetStatement(
							ast.NewIdentifier("foo"),
							ast.NewBoolean(true),
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
				let true;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{},
				}
				expectedErrors := []error{
					parser.ErrUnexpectedTokenType,
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
				let foo true;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{},
				}
				expectedErrors := []error{
					parser.ErrUnexpectedTokenType,
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
			It("correct program", func() {
				text = `
				return 5;
				return 10;
				return 838383;
				return true;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewReturnStatement(ast.NewInteger("5", 5)),
						ast.NewReturnStatement(ast.NewInteger("10", 10)),
						ast.NewReturnStatement(ast.NewInteger("838383", 838383)),
						ast.NewReturnStatement(ast.NewBoolean(true)),
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
