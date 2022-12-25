package lexer

import (
	"testing"

	"github.com/sammyoina/boa/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test %d expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tets %d expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
