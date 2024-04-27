package lexer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Aden-Q/monkey/internal/lexer"
	"github.com/Aden-Q/monkey/internal/token"
)

var _ = Describe("Lexer", func() {

	Describe("NextToken test", func() {
		Context("valid input", func() {
			It("should equal", func() {
				input := `=+(){},;`
				expected_tokens := []*token.Token{
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
	})

})
