package token

type TokenType int

const (
	ILLEGAL TokenType = 1
	EOF     TokenType = 2

	IDENTIFIER TokenType = 3
	VAR_TYPE   TokenType = 4
	INT        TokenType = 5

	ASSIGN TokenType = 6
	PLUS   TokenType = 7

	COMMA  TokenType = 8
	LPAREN TokenType = 9
	RPAREN TokenType = 10
	LBRACE TokenType = 11
	RBRACE TokenType = 12

	FUNCTION TokenType = 13
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def": FUNCTION,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
