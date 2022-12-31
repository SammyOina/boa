package parser

import (
	"fmt"
	"testing"

	"github.com/sammyoina/boa/ast"
	"github.com/sammyoina/boa/lexer"
)

func TestAssignStatements(t *testing.T) {
	input := `
	x = 5
	y = 3
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("Parse returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("wanted 2 statements got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	assignStatement, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s not assign statement got %T", s)
		return false
	}
	/*if assignStatement.Name.Value != name {
		t.Errorf("wanted %s got %s", name, assignStatement.Name.Value)
		return false
	}*/
	/*if assignStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s got %s", name, assignStatement.Name.Token.Literal)
		return false
	}*/
	fmt.Println(assignStatement.Name.Value, assignStatement.Value)
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5
	return 11001
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("Required 2 program statements, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement, got %T", returnStmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return' gor %q", returnStmt.TokenLiteral())
		}
	}

}

func TestIdentifierExpression(t *testing.T) {
	input := "x"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("not enough statements got, %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement not of type ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("expression not of type Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "x" {
		t.Errorf("value not 'x', got %s", ident.Value)
	}
	if ident.TokenLiteral() != "x" {
		t.Errorf("identifier literal not x, got %s", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("not enough statements wanted 1 got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statment not expression statment got %T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression not integer, got %T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal value not 5, got %d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not 5 got %s", literal.TokenLiteral())
	}

}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, test := range prefixTests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("not engough program statements got %d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program statement not expression, got %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expression not prefix expression, got %T", stmt.Expression)
		}
		if exp.Operator != test.operator {
			t.Fatalf("Wrong Operator, wanted %s, got %s", test.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, test.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("not integer literal, got %T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("wanted %d, got %d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d, got %s", value, integ.TokenLiteral())
		return false
	}

	return true
}
