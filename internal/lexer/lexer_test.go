package lexer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aden-q/monkey/internal/lexer"
	"github.com/aden-q/monkey/internal/token"
)

var _ = Describe("Lexer", func() {
	var (
		text string
		l    lexer.Lexer
	)

	BeforeEach(func() {
		l = lexer.New()
	})

	Describe("Lexer", func() {
		Context("Read", func() {
			It("can read simple text", func() {
				text = `something`
				Expect(l.Read(text)).To(Equal(len(text)))
			})
		})

		Context("NextToken", func() {
			It("can parse simple text", func() {
				text = `=+(){},;`
				expectedTokens := []token.Token{
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

				Expect(l.Read(text)).To(Equal(len(text)))

				for _, expectedToken := range expectedTokens {
					token := l.NextToken()
					Expect(token).To(Equal(expectedToken))
				}
			})
		})

		Context("code snippet", func() {
			It("can parse complex text", func() {
				text =
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
					};
					
					10 == 10;
					10 != 9;
					9 <= 10;
					"foo";
					"foo bar";
					`
				expectedTokens := []token.Token{
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
						Type:    token.FUNC,
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
					{
						Type:    token.INT,
						Literal: "10",
					},
					{
						Type:    token.EQ,
						Literal: "==",
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
						Type:    token.INT,
						Literal: "10",
					},
					{
						Type:    token.NOT_EQ,
						Literal: "!=",
					},
					{
						Type:    token.INT,
						Literal: "9",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.INT,
						Literal: "9",
					},
					{
						Type:    token.LTE,
						Literal: "<=",
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
						Type:    token.STRING,
						Literal: "foo",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.STRING,
						Literal: "foo bar",
					},
					{
						Type:    token.SEMICOLON,
						Literal: ";",
					},
					{
						Type:    token.EOF,
						Literal: "eof",
					},
				}

				Expect(l.Read(text)).To(Equal(len(text)))

				for _, expectedToken := range expectedTokens {
					token := l.NextToken()
					Expect(token).To(Equal(expectedToken))
				}
			})
		})
	})

})
