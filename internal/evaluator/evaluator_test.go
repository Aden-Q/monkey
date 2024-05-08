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
			It("integer expressions", func() {
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
		})

		Context("boolean object", func() {
			It("boolean expressions", func() {
				text = `
				true;
				`
				expectedObject := object.NewBoolean(true)
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
