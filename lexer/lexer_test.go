package lexer

import (
	"fmt"
	"testing"

	"github.com/sammyoina/boa/token"
)

func TestNextToken(t *testing.T) {
	input := `a = 5
	b  = 6
	def test(x,y) {
		return x + y
	}
	res = test(a,b)
	{"foo":"bar"}
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENTIFIER, "a"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NewLine, "\n"},
		{token.IDENTIFIER, "b"},
		{token.ASSIGN, "="},
		{token.INT, "6"},
		{token.NewLine, "\n"},
		{token.FUNCTION, "def"},
		{token.IDENTIFIER, "test"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NewLine, "\n"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.NewLine, "\n"},
		{token.RBRACE, "}"},
		{token.NewLine, "\n"},
		{token.IDENTIFIER, "res"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "test"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "a"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "b"},
		{token.RPAREN, ")"},
		{token.NewLine, "\n"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.NewLine, "\n"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("test %d expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tets %d expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
