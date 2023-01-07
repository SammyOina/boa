package token

type TokenType int

const (
	_ TokenType = iota
	ILLEGAL
	EOF

	IDENTIFIER
	//VAR_TYPE
	INT
	STRING

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
	LBRACKET
	RBRACKET
	COLON

	FUNCTION
	TRUE
	FALSE
	IF
	ELSE
	RETURN

	NewLine
)

var TokenTypeString = map[TokenType]string{
	ILLEGAL:    "illegal",
	EOF:        "EOF",
	IDENTIFIER: "IDENTIFIER",
	//VAR_TYPE:     "VAR_TYPE",
	INT:          "Integer",
	ASSIGN:       "=",
	PLUS:         "+",
	MINUS:        "-",
	BANG:         "!",
	ASTERISK:     "*",
	SLASH:        "/",
	LESS_THAN:    "<",
	GREATER_THAN: ">",
	EQUAL_TO:     "==",
	NOT_EQUAL_TO: "!=",
	COMMA:        ",",
	LPAREN:       "(",
	RPAREN:       ")",
	LBRACE:       "{",
	RBRACE:       "}",
	LBRACKET:     "[",
	RBRACKET:     "]",
	FUNCTION:     "FUNCTION",
	TRUE:         "TRUE",
	FALSE:        "FALSE",
	IF:           "IF",
	ELSE:         "ELSE",
	RETURN:       "RETURN",
	NewLine:      "NewLine",
	STRING:       "\"",
	COLON:        ":",
}

var TokenTypeByte = map[TokenType]byte{
	ASSIGN:       '=',
	PLUS:         '+',
	MINUS:        '-',
	BANG:         '!',
	ASTERISK:     '*',
	SLASH:        '/',
	LESS_THAN:    '<',
	GREATER_THAN: '>',
	COMMA:        ',',
	LPAREN:       '(',
	RPAREN:       ')',
	LBRACE:       '{',
	RBRACE:       '}',
	NewLine:      '\n',
	STRING:       '"',
	LBRACKET:     '[',
	RBRACKET:     ']',
	EOF:          0,
	COLON:        ':',
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
