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
		Context("expressions", func() {
			It("identifier expressions", func() {
				text = `
				foo;
				bar;
				cs;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewIdentifierExpression("foo")),
						ast.NewExpressionStatement(ast.NewIdentifierExpression("bar")),
						ast.NewExpressionStatement(ast.NewIdentifierExpression("cs")),
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
						ast.NewExpressionStatement(ast.NewIntegerExpression("5", 5)),
						ast.NewExpressionStatement(ast.NewIntegerExpression("10", 10)),
						ast.NewExpressionStatement(ast.NewIntegerExpression("838383", 838383)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("boolean expressions", func() {
				text = `
				true;
				false;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewBooleanExpression(true)),
						ast.NewExpressionStatement(ast.NewBooleanExpression(false)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("string expressions", func() {
				text = `
				"foo";
				"foo bar";
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewStringExpression("foo")),
						ast.NewExpressionStatement(ast.NewStringExpression("foo bar")),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("array expressions", func() {
				text = `
				[1, 2 * 2, 3 + 3];
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewArrayExpression(ast.NewIntegerExpression("1", 1), ast.NewInfixExpression("*", ast.NewIntegerExpression("2", 2), ast.NewIntegerExpression("2", 2)), ast.NewInfixExpression("+", ast.NewIntegerExpression("3", 3), ast.NewIntegerExpression("3", 3)))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("hash expressions", func() {
				text = `
				{"foo": 5, "bar": "hi", "cs": true};
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewHashExpression(map[ast.Expression]ast.Expression{
							ast.NewStringExpression("foo"): ast.NewIntegerExpression("5", 5),
							ast.NewStringExpression("bar"): ast.NewStringExpression("hi"),
							ast.NewStringExpression("cs"):  ast.NewBooleanExpression(true),
						}),
						),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				// FIXME: hash expression is not comparable, order is random
				Expect(len(program.Statements)).To(Equal(len(expectedProgram.Statements)))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("empty hash expressions", func() {
				text = `
				{};
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewHashExpression(map[ast.Expression]ast.Expression{})),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("index expressions", func() {
				text = `
				myArray[1+1];
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewIndexExpression(ast.NewIdentifierExpression("myArray"), ast.NewInfixExpression("+", ast.NewIntegerExpression("1", 1), ast.NewIntegerExpression("1", 1)))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("if expressions", func() {
				text = `
				if (x < y) { x; };
				if (x < y) { x; } else { y; };
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewIfExpression(
							ast.NewInfixExpression("<", ast.NewIdentifierExpression("x"), ast.NewIdentifierExpression("y")),
							ast.NewBlockStatement(
								ast.NewExpressionStatement(ast.NewIdentifierExpression("x")),
							),
							nil,
						)),
						ast.NewExpressionStatement(ast.NewIfExpression(
							ast.NewInfixExpression("<", ast.NewIdentifierExpression("x"), ast.NewIdentifierExpression("y")),
							ast.NewBlockStatement(ast.NewExpressionStatement(ast.NewIdentifierExpression("x"))),
							ast.NewBlockStatement(
								ast.NewExpressionStatement(ast.NewIdentifierExpression("y")),
							),
						)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("func expressions", func() {
				text = `
				fn() {};
				fn(x) { 1; };
				fn(x, y) { x + y; };
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewFuncExpression(
							[]*ast.IdentifierExpression{},
							ast.NewBlockStatement(),
						)),
						ast.NewExpressionStatement(ast.NewFuncExpression(
							[]*ast.IdentifierExpression{
								ast.NewIdentifierExpression("x"),
							},
							ast.NewBlockStatement(
								ast.NewExpressionStatement(ast.NewIntegerExpression("1", 1))),
						)),
						ast.NewExpressionStatement(ast.NewFuncExpression(
							[]*ast.IdentifierExpression{
								ast.NewIdentifierExpression("x"),
								ast.NewIdentifierExpression("y"),
							},
							ast.NewBlockStatement(
								ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifierExpression("x"), ast.NewIdentifierExpression("y"))),
							),
						)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("call expressions", func() {
				text = `
				add(1, 2 * 3, 4 + 5);
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(ast.NewCallExpression(
							ast.NewIdentifierExpression("add"),
							[]ast.Expression{
								ast.NewIntegerExpression("1", 1),
								ast.NewInfixExpression("*", ast.NewIntegerExpression("2", 2), ast.NewIntegerExpression("3", 3)),
								ast.NewInfixExpression("+", ast.NewIntegerExpression("4", 4), ast.NewIntegerExpression("5", 5)),
							},
						)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("prefix expressions", func() {
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
						ast.NewExpressionStatement(ast.NewPrefixExpression("-", ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewIntegerExpression("10", 10))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("-", ast.NewIdentifierExpression("foo"))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewIdentifierExpression("bar"))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewBooleanExpression(true))),
						ast.NewExpressionStatement(ast.NewPrefixExpression("!", ast.NewBooleanExpression(false))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("infix expressions", func() {
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
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("-", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("*", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("/", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression(">", ast.NewIntegerExpression("6", 6), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression(">=", ast.NewIntegerExpression("6", 6), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("<", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("<=", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("==", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("!=", ast.NewIntegerExpression("5", 5), ast.NewIntegerExpression("6", 6))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifierExpression("foo"), ast.NewIdentifierExpression("bar"))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifierExpression("foo"), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifierExpression("bar"), ast.NewIntegerExpression("5", 5))),
						ast.NewExpressionStatement(ast.NewInfixExpression("*", ast.NewPrefixExpression("-", ast.NewIdentifierExpression("a")), ast.NewIdentifierExpression("b"))),
						ast.NewExpressionStatement(ast.NewInfixExpression("==", ast.NewBooleanExpression(true), ast.NewBooleanExpression(true))),
						ast.NewExpressionStatement(ast.NewInfixExpression("!=", ast.NewBooleanExpression(true), ast.NewBooleanExpression(false))),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})

			It("expression string match", func() {
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
					`a + add(b * c) +d;`,
					`add(a + b + c * d / f + g);`,
					`a * [1, 2, 3, 4][b * c] * d;`,
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
					`((a + add((b * c))) + d)`,
					`add((((a + b) + ((c * d) / f)) + g))`,
					`((a * ([1, 2, 3, 4][(b * c)])) * d)`,
				}
				expectedErrors := []error{}

				for idx := range texts {
					program, errs = p.ParseProgram(texts[idx])
					Expect(program.String()).To(Equal(expectedStrings[idx]))
				}

				Expect(errs).To(Equal(expectedErrors))
			})
		})

		Context("let statements", func() {
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
							ast.NewIdentifierExpression("x"),
							ast.NewIntegerExpression("5", 5),
						),
						ast.NewLetStatement(
							ast.NewIdentifierExpression("y"),
							ast.NewIntegerExpression("10", 10),
						),
						ast.NewLetStatement(
							ast.NewIdentifierExpression("foo"),
							ast.NewIntegerExpression("838383", 838383),
						),
						ast.NewLetStatement(
							ast.NewIdentifierExpression("foo"),
							ast.NewBooleanExpression(true),
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

		Context("return statements", func() {
			It("correct program", func() {
				text = `
				return 5;
				return 10;
				return 838383;
				return true;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewReturnStatement(ast.NewIntegerExpression("5", 5)),
						ast.NewReturnStatement(ast.NewIntegerExpression("10", 10)),
						ast.NewReturnStatement(ast.NewIntegerExpression("838383", 838383)),
						ast.NewReturnStatement(ast.NewBooleanExpression(true)),
					},
				}
				expectedErrors := []error{}

				program, errs = p.ParseProgram(text)
				Expect(program).To(Equal(expectedProgram))
				Expect(errs).To(Equal(expectedErrors))
			})
		})

		Context("special statements", func() {
			It("empty statement", func() {
				text = `
				;
				`
				expectedProgram := &ast.Program{
					Statements: []ast.Statement{
						ast.NewExpressionStatement(nil),
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
