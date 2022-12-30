package token

type TokenType int

const (
	ILLEGAL TokenType = 1
	EOF     TokenType = 2

	IDENTIFIER TokenType = 3
	VAR_TYPE   TokenType = 4
	INT        TokenType = 5

	ASSIGN   TokenType = 6
	PLUS     TokenType = 7
	MINUS    TokenType = 14
	BANG     TokenType = 15
	ASTERISK TokenType = 16
	SLASH    TokenType = 17

	LESS_THAN    TokenType = 18
	GREATER_THAN TokenType = 19
	EQUAL_TO     TokenType = 25
	NOT_EQUAL_TO TokenType = 26

	COMMA  TokenType = 8
	LPAREN TokenType = 9
	RPAREN TokenType = 10
	LBRACE TokenType = 11
	RBRACE TokenType = 12

	FUNCTION TokenType = 13
	TRUE     TokenType = 20
	FALSE    TokenType = 21
	IF       TokenType = 22
	ELSE     TokenType = 23
	RETURN   TokenType = 24

	NewLine TokenType = 27
)

var TokenTypeString = map[TokenType]string{
	ILLEGAL:      "illegal",
	EOF:          "EOF",
	IDENTIFIER:   "IDENTIFIER",
	VAR_TYPE:     "VAR_TYPE",
	INT:          "INT",
	ASSIGN:       "ASSIGN",
	PLUS:         "PLUS",
	MINUS:        "MINUS",
	BANG:         "BANG",
	ASTERISK:     "ASTERISK",
	SLASH:        "SLASH",
	LESS_THAN:    "LESS_THAN",
	GREATER_THAN: "GREATER_THAN",
	EQUAL_TO:     "EQUAL_TO",
	NOT_EQUAL_TO: "NOT_EQUAL_TO",
	COMMA:        "COMMA",
	LPAREN:       "LPAREN",
	RPAREN:       "RPAREN",
	LBRACE:       "LBRACE",
	RBRACE:       "RBRACE",
	FUNCTION:     "FUNCTION",
	TRUE:         "TRUE",
	FALSE:        "FALSE",
	IF:           "IF",
	ELSE:         "ELSE",
	RETURN:       "RETURN",
	NewLine:      "NewLine",
}

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def":    FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
