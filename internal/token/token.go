package token

import "strconv"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 123456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywordTable = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

var operatorTable = map[string]TokenType{
	"=": ASSIGN,
	"+": PLUS,
}

var delimeterTable = map[string]TokenType{
	",": COMMA,
	";": SEMICOLON,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func LookupTokenType(literal string) TokenType {
	if tokType, ok := keywordTable[literal]; ok {
		return tokType
	}

	if tokType, ok := operatorTable[literal]; ok {
		return tokType
	}

	if tokType, ok := delimeterTable[literal]; ok {
		return tokType
	}

	if _, err := strconv.Atoi(literal); err == nil {
		return INT
	}

	// as for now, treat the default type as identifier
	return IDENT
}
