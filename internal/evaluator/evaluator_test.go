package evaluator_test

import (
	"github.com/aden-q/monkey/internal/ast"
	"github.com/aden-q/monkey/internal/evaluator"
	"github.com/aden-q/monkey/internal/lexer"
	"github.com/aden-q/monkey/internal/object"
	"github.com/aden-q/monkey/internal/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Evaluator", func() {
	var (
		text    string
		l       lexer.Lexer
		p       parser.Parser
		e       evaluator.Evaluator
		program *ast.Program
		errs    []error
	)

	BeforeEach(func() {
		l = lexer.New()
		p = parser.New(l)
		e = evaluator.New(object.NewEnvironment())
	})

	Describe("Eval", func() {
		Context("integer object", func() {
			It("integer expression", func() {
				text = `
				5;
				`
				expectedObject := object.NewInteger(5)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix minus operator expression", func() {
				text = `
				-5;
				`
				expectedObject := object.NewInteger(-5)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix plus operator expression", func() {
				text = `
				5 + 5;
				`
				expectedObject := object.NewInteger(10)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix minus operator expression", func() {
				text = `
				5 - 5;
				`
				expectedObject := object.NewInteger(0)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix multiply operator expression", func() {
				text = `
				5 * 5;
				`
				expectedObject := object.NewInteger(25)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix division operator expression", func() {
				text = `
				10 / 5;
				`
				expectedObject := object.NewInteger(2)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix mix-operator expression", func() {
				text = `
				5 + 5 - 2 + 10 * 3 / 5;
				`
				expectedObject := object.NewInteger(14)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})
		})

		Context("boolean object", func() {
			It("boolean expression", func() {
				text = `
				true;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!true;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!false;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!5;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!!true;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!!false;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("prefix bang boolean expression", func() {
				text = `
				!!5;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 < 2;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 <= 2;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 < 1;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 <= 1;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 > 2;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 >= 2;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 > 1;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 >= 1;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 == 1;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 == 1;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				2 != 1;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				1 != 1;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				true == true;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				false == false;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				true == false;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				true != true;
				`
				expectedObject := object.FALSE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix comparision expression", func() {
				text = `
				true != false;
				`
				expectedObject := object.TRUE
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})
		})

		Context("string object", func() {
			It("string expression", func() {
				text = `
				"hello world";
				`
				expectedObject := object.NewString("hello world")
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("infix string concatenation", func() {
				text = `
				"hello" + " " + "world!";
				`
				expectedObject := object.NewString("hello world!")
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})
		})

		Context("array object", func() {
			It("array expression", func() {
				text = `
				[1, 2 * 2, true];
				`
				expectedObject := object.NewArray(object.NewInteger(1), object.NewInteger(4), object.NewBoolean(true))
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("index expression", func() {
				text = `
				[1, 2 * 2, true][0];
				`
				expectedObject := object.NewInteger(1)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("index expression", func() {
				text = `
				[1, 2 * 2, true][2];
				`
				expectedObject := object.NewBoolean(true)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("index expression", func() {
				text = `
				let a = [1, 2 * 2, true];
				a[1];
				`
				expectedObject := object.NewInteger(4)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("index expression, index out of range", func() {
				text = `
				let a = [1, 2 * 2, true];
				a[3];
				`
				expectedObject := object.NIL
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).To(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})
		})

		Context("if conditionals", func() {
			It("if condition is truthy", func() {
				text = `
				if (true) {
					10;
				};
				`
				expectedObject := object.NewInteger(10)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("if condition is truthy", func() {
				text = `
				if (5) {
					10;
				};
				`
				expectedObject := object.NewInteger(10)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("if condition is truthy", func() {
				text = `
				if (1 < 2) {
					10;
				};
				`
				expectedObject := object.NewInteger(10)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("if condition is false", func() {
				text = `
				if (false) {
					10;
				};
				`
				expectedObject := object.NIL
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})
		})

		Context("return statements", func() {
			It("return an integer", func() {
				text = `
				return 10;
				`
				expectedObject := object.NewReturnValue(object.NewInteger(10))
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("return an integer", func() {
				text = `
				9;
				false;
				return 10;
				5;
				true;
				`
				expectedObject := object.NewReturnValue(object.NewInteger(10))
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("return an integer with if condition", func() {
				text = `
				if (10 > 1) {
					return 10;
				};

				return 5;
				`
				expectedObject := object.NewReturnValue(object.NewInteger(10))
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("return an integer with nested if conditions", func() {
				text = `
				if (10 > 1) {
					if (10 > 1) {
						return 10;
					};

					return 8;
				};
				
				return 5;
				`
				expectedObject := object.NewReturnValue(object.NewInteger(10))
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			It("return an boolean with if condition", func() {
				text = `
				if (10 > 1) {
					return 10 > 1;
				};
				
				return false;
				`
				expectedObject := object.NewReturnValue(object.TRUE)
				expectedErrors := []error{}

				// parse the program
				program, errs = p.ParseProgram(text)
				Expect(errs).To(Equal(expectedErrors))

				// evaluate the AST tree
				obj, err := e.Eval(program)
				Expect(err).ToNot(HaveOccurred())
				Expect(obj).To(Equal(expectedObject))
			})

			Context("errors", func() {
				It("can detect errors", func() {
					text = `
					5 + true;
					`
					expectedObject := object.NIL
					expectedParseErrors := []error{}
					expectedEvaluateError := evaluator.ErrUnexpectedObjectType

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).To(Equal(expectedEvaluateError))
					Expect(obj).To(Equal(expectedObject))
				})
			})

			Context("errors", func() {
				It("can early terminate when there's an error", func() {
					text = `
					5 + true;
					10;
					`
					expectedObject := object.NIL
					expectedParseErrors := []error{}
					expectedEvaluateError := evaluator.ErrUnexpectedObjectType

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).To(Equal(expectedEvaluateError))
					Expect(obj).To(Equal(expectedObject))
				})
			})

			Context("let statements", func() {
				It("successful binding", func() {
					text = `
					let a = 5 * 5;
					a;
					`
					expectedObject := object.NewInteger(25)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("unbound identifier", func() {
					text = `
					foobar;
					`
					expectedObject := object.NIL
					expectedParseErrors := []error{}
					expectedEvaluateError := evaluator.ErrIdentifierNotFound

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).To(Equal(expectedEvaluateError))
					Expect(obj).To(Equal(expectedObject))
				})
			})

			Context("function expressions", func() {
				It("func", func() {
					text = `
					fn(x) { x + 2; };
					`
					expectedObject := object.NewFunc(
						[]*ast.IdentifierExpression{
							ast.NewIdentifierExpression("x"),
						},
						ast.NewBlockStatement(ast.NewExpressionStatement(ast.NewInfixExpression("+", ast.NewIdentifierExpression("x"), ast.NewIntegerExpression("2", 2)))),
						object.NewEnvironment(),
					)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})
			})

			Context("call expressions", func() {
				It("func with 1 parameter", func() {
					text = `
					let a = fn(x) { x + 2; };
					a(5);
					`
					expectedObject := object.NewInteger(7)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("func with 2 parameters", func() {
					text = `
					let a = fn(x, y) { x * y + 2; };
					a(5, 6);
					`
					expectedObject := object.NewInteger(32)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("func with 2 parameters", func() {
					text = `
				 	fn(x, y) { x * y + 2; } (5, 6);
					`
					expectedObject := object.NewInteger(32)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})
			})

			Context("builtin functions", func() {
				It("length of an empty string", func() {
					text = `
				 	len("");
					`
					expectedObject := object.NewInteger(0)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("length of a non-empty string", func() {
					text = `
						 len("hello world");
						`
					expectedObject := object.NewInteger(11)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("length on integer unsupported", func() {
					text = `
						 len(1);
						`
					expectedObject := object.NIL
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).To(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})

				It("length of an array", func() {
					text = `
						 len([1,2,true]);
						`
					expectedObject := object.NewInteger(3)
					expectedParseErrors := []error{}

					// parse the program
					program, errs = p.ParseProgram(text)
					Expect(errs).To(Equal(expectedParseErrors))

					// evaluate the AST tree
					obj, err := e.Eval(program)
					Expect(err).ToNot(HaveOccurred())
					Expect(obj).To(Equal(expectedObject))
				})
			})
		})
	})
})
