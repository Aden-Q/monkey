package ast_test

import (
	"github.com/aden-q/monkey/internal/ast"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ast", func() {
	var (
		identifierFoo   *ast.IdentifierExpression
		identifierBar   *ast.IdentifierExpression
		letStatement    *ast.LetStatement
		returnStatement *ast.ReturnStatement
		program         *ast.Program
	)

	BeforeEach(func() {
		// let foo = bar;
		identifierFoo = ast.NewIdentifierExpression("foo")
		Expect(identifierFoo).ToNot(BeNil())
		identifierBar = ast.NewIdentifierExpression("bar")
		Expect(identifierBar).ToNot(BeNil())
		letStatement = ast.NewLetStatement(identifierFoo, identifierBar)
		Expect(letStatement).ToNot(BeNil())
		program = ast.NewProgram(letStatement)
		Expect(program).ToNot(BeNil())
		returnStatement = ast.NewReturnStatement(nil)
	})

	Describe("TokenLiteral", func() {
		It("Program TokenLiteral", func() {
			Expect(program.TokenLiteral()).To(Equal("let"))
		})

		It("Identifier TokenLiteral", func() {
			Expect(identifierFoo.TokenLiteral()).To(Equal("foo"))
			Expect(identifierBar.TokenLiteral()).To(Equal("bar"))
		})

		It("LetStatement TokenLiteral", func() {
			Expect(letStatement.TokenLiteral()).To(Equal("let"))
		})

		It("ReturnStatement TokenLiteral", func() {
			Expect(returnStatement.TokenLiteral()).To(Equal("return"))
		})

		It("ExpressionStatement TokenLiteral", func() {
		})
	})

	Describe("String", func() {
		Context("Program as a string", func() {
			It("string output matches the program", func() {
				Expect(program.String()).To(Equal("let foo = bar;"))
			})
		})

		Context("Identifier as a string", func() {
			It("string output matches the identifier", func() {
				Expect(identifierFoo.String()).To(Equal("foo"))
				Expect(identifierBar.String()).To(Equal("bar"))
			})
		})

		Context("LetStatement as a string", func() {
			It("string output matches the let statement", func() {
				Expect(letStatement.String()).To(Equal("let foo = bar;"))
			})
		})

		Context("ReturnStatement as a string", func() {
			It("string output matches the return statement", func() {
			})
		})

		Context("ExpressionStatement as a string", func() {
			It("string output matches the expression statement", func() {
			})
		})
	})
})
