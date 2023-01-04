package evaluator

import (
	"testing"

	"github.com/sammyoina/boa/lexer"
	"github.com/sammyoina/boa/object"
	"github.com/sammyoina/boa/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnv()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object not integer got, %T", obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("wanted %d, got %d", expected, result.Value)
		return false
	}
	return true
}
