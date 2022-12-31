package token

type TokenType int

const (
	_ TokenType = iota
	ILLEGAL
	EOF

	IDENTIFIER
	VAR_TYPE
	INT

	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	LESS_THAN
	GREATER_THAN
	EQUAL_TO
	NOT_EQUAL_TO

	COMMA
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	TRUE
	FALSE
	IF
	ELSE
	RETURN

	NewLine
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
