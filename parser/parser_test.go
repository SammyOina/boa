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
