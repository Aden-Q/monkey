package token

import "strconv"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 123456
	STRING = "STRING"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	LTE      = "<="
	GT       = ">"
	GTE      = ">="
	EQ       = "=="
	NOT_EQ   = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// keywords
	FUNC   = "FUNC"
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywordTable = map[string]TokenType{
	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

var operatorTable = map[string]TokenType{
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"!":  BANG,
	"*":  ASTERISK,
	"/":  SLASH,
	"<":  LT,
	"<=": LTE,
	">":  GT,
	">=": GTE,
	"==": EQ,
	"!=": NOT_EQ,
}

var delimeterTable = map[string]TokenType{
	",": COMMA,
	";": SEMICOLON,
	":": COLON,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
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
	// TODO: more checks to separate between IDENT and ILLEGAL
	return IDENT
}
