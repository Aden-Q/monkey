package token

// precedence levels
const (
	LOWEST      = iota
	EQUALS      // ==
	LESSGREATER // >, >=, <, <=
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // fn(x) { return x + 1; } (1);
)

// token type to precedence maping
var (
	precedences = map[TokenType]int{
		EQ:       EQUALS,
		NOT_EQ:   EQUALS,
		LT:       LESSGREATER,
		LTE:      LESSGREATER,
		GT:       LESSGREATER,
		GTE:      LESSGREATER,
		PLUS:     SUM,
		MINUS:    SUM,
		SLASH:    PRODUCT,
		ASTERISK: PRODUCT,
		BANG:     PREFIX,
		LPAREN:   CALL,
	}
)

func GetPrecedence(tokenType TokenType) int {
	if p, ok := precedences[tokenType]; ok {
		return p
	}

	return LOWEST
}
