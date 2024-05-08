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
		e = evaluator.New()
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
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
				Expect(err).To(BeNil())
				Expect(obj).To(Equal(expectedObject))
			})
		})
	})
})
