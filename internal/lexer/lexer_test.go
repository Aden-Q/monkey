package lexer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Aden-Q/monkey/internal/lexer"
	"github.com/Aden-Q/monkey/internal/token"
)

var _ = Describe("Lexer", func() {

	Describe("NextToken test", func() {
		Context("simple lexing test", func() {
			It("should equal", func() {
				input := `=+(){},;`
				expected_tokens := []token.Token{
					{
						Type:    token.ASSIGN,
						Literal: "=",
					},
					{
						Type:    token.PLUS,
						Literal: "+",
					},
					{
						Type:    token.LPAREN,
						Literal: "(",
					},
					{
						Type:    token.RPAREN,
						Literal: ")",
					},
					{
						Type:    token.LBRACE,
						Literal: "{",
					},
					{
						Type:    token.RBRACE,
						Literal: "}",
					},
					{
						Type:    token.COMMA,
						Literal: ",",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
				}

				l := lexer.New(input)

				for _, expected_token := range expected_tokens {
					token, ok := l.NextToken()
					Expect(ok).To(Equal(true))
					Expect(token).To(Equal(expected_token))
				}
			})
		})

		Context("complex lexing test", func() {
			It("should equal", func() {
				input :=
					`let five = 5;
					let ten = 10;

					let add = fn(x, y) {
						x + y;
					};
					
					let result = add(five, ten);

					!-/*5;
					5 < 10 > 5;

					if (5 < 10) {
						return true
					} else {
						return false
					};`
				expected_tokens := []token.Token{
					{
						Type:    token.LET,
						Literal: "let",
					},
					{
						Type:    token.IDENT,
						Literal: "five",
					},
					{
						Type:    token.ASSIGN,
						Literal: "=",
					},
					{
						Type:    token.INT,
						Literal: "5",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.LET,
						Literal: "let",
					},
					{
						Type:    token.IDENT,
						Literal: "ten",
					},
					{
						Type:    token.ASSIGN,
						Literal: "=",
					},
					{
						Type:    token.INT,
						Literal: "10",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.LET,
						Literal: "let",
					},
					{
						Type:    token.IDENT,
						Literal: "add",
					},
					{
						Type:    token.ASSIGN,
						Literal: "=",
					},
					{
						Type:    token.FUNCTION,
						Literal: "fn",
					},
					{
						Type:    token.LPAREN,
						Literal: "(",
					},
					{
						Type:    token.IDENT,
						Literal: "x",
					},
					{
						Type:    token.COMMA,
						Literal: ",",
					},
					{
						Type:    token.IDENT,
						Literal: "y",
					},
					{
						Type:    token.RPAREN,
						Literal: ")",
					},
					{
						Type:    token.LBRACE,
						Literal: "{",
					},
					{
						Type:    token.IDENT,
						Literal: "x",
					},
					{
						Type:    token.PLUS,
						Literal: "+",
					},
					{
						Type:    token.IDENT,
						Literal: "y",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.RBRACE,
						Literal: "}",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},

					{
						Type:    token.LET,
						Literal: "let",
					},
					{
						Type:    token.IDENT,
						Literal: "result",
					},
					{
						Type:    token.ASSIGN,
						Literal: "=",
					},
					{
						Type:    token.IDENT,
						Literal: "add",
					},
					{
						Type:    token.LPAREN,
						Literal: "(",
					},
					{
						Type:    token.IDENT,
						Literal: "five",
					},
					{
						Type:    token.COMMA,
						Literal: ",",
					},
					{
						Type:    token.IDENT,
						Literal: "ten",
					},
					{
						Type:    token.RPAREN,
						Literal: ")",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.BANG,
						Literal: "!",
					},
					{
						Type:    token.MINUS,
						Literal: "-",
					},
					{
						Type:    token.SLASH,
						Literal: "/",
					},
					{
						Type:    token.ASTERISK,
						Literal: "*",
					},
					{
						Type:    token.INT,
						Literal: "5",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.INT,
						Literal: "5",
					},
					{
						Type:    token.LT,
						Literal: "<",
					},
					{
						Type:    token.INT,
						Literal: "10",
					},
					{
						Type:    token.GT,
						Literal: ">",
					},
					{
						Type:    token.INT,
						Literal: "5",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.IF,
						Literal: "if",
					},
					{
						Type:    token.LPAREN,
						Literal: "(",
					},
					{
						Type:    token.INT,
						Literal: "5",
					},
					{
						Type:    token.LT,
						Literal: "<",
					},
					{
						Type:    token.INT,
						Literal: "10",
					},
					{
						Type:    token.RPAREN,
						Literal: ")",
					},
					{
						Type:    token.LBRACE,
						Literal: "{",
					},
					{
						Type:    token.RETURN,
						Literal: "return",
					},
					{
						Type:    token.TRUE,
						Literal: "true",
					},
					{
						Type:    token.RBRACE,
						Literal: "}",
					},
					{
						Type:    token.ELSE,
						Literal: "else",
					},
					{
						Type:    token.LBRACE,
						Literal: "{",
					},
					{
						Type:    token.RETURN,
						Literal: "return",
					},
					{
						Type:    token.FALSE,
						Literal: "false",
					},
					{
						Type:    token.RBRACE,
						Literal: "}",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
				}

				l := lexer.New(input)

				for _, expected_token := range expected_tokens {
					token, ok := l.NextToken()
					Expect(ok).To(Equal(true))
					Expect(token).To(Equal(expected_token))
				}
			})
		})
	})

})
