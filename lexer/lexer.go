package lexer

import "github.com/sammyoina/boa/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case token.TokenTypeByte[token.ASSIGN]:
		if l.peekChar() == '=' {
			//ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQUAL_TO, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case token.TokenTypeByte[token.LPAREN]:
		tok = newToken(token.LPAREN, l.ch)
	case token.TokenTypeByte[token.RPAREN]:
		tok = newToken(token.RPAREN, l.ch)
	case token.TokenTypeByte[token.COMMA]:
		tok = newToken(token.COMMA, l.ch)
	case token.TokenTypeByte[token.PLUS]:
		tok = newToken(token.PLUS, l.ch)
	case token.TokenTypeByte[token.LBRACE]:
		tok = newToken(token.LBRACE, l.ch)
	case token.TokenTypeByte[token.RBRACE]:
		tok = newToken(token.RBRACE, l.ch)
	case token.TokenTypeByte[token.MINUS]:
		tok = newToken(token.MINUS, l.ch)
	case token.TokenTypeByte[token.BANG]:
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQUAL_TO, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case token.TokenTypeByte[token.LESS_THAN]:
		tok = newToken(token.LESS_THAN, l.ch)
	case token.TokenTypeByte[token.GREATER_THAN]:
		tok = newToken(token.GREATER_THAN, l.ch)
	case token.TokenTypeByte[token.ASTERISK]:
		tok = newToken(token.ASTERISK, l.ch)
	case token.TokenTypeByte[token.SLASH]:
		tok = newToken(token.SLASH, l.ch)
	case token.TokenTypeByte[token.NewLine]:
		tok = newToken(token.NewLine, l.ch)
	case token.TokenTypeByte[token.STRING]:
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case token.TokenTypeByte[token.EOF]:
		tok.Literal = ""
		tok.Type = token.EOF
	case token.TokenTypeByte[token.LBRACKET]:
		tok = newToken(token.LBRACKET, l.ch)
	case token.TokenTypeByte[token.RBRACKET]:
		tok = newToken(token.RBRACKET, l.ch)
	case token.TokenTypeByte[token.COLON]:
		tok = newToken(token.COLON, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' /*|| l.ch == '\n'*/ || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}
