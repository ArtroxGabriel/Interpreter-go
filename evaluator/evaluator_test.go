package evaluator_test

import (
	"testing"

	"github.com/ArtroxGabriel/interpreter/evaluator"
	"github.com/ArtroxGabriel/interpreter/lexer"
	"github.com/ArtroxGabriel/interpreter/object"
	"github.com/ArtroxGabriel/interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int64
	}{
		{
			input:    "5",
			expected: 5,
		},
		{
			input:    "10",
			expected: 10,
		},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testIntegerObject(t, evaluated, tC.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			input:    "true",
			expected: true,
		},
		{
			input:    "false",
			expected: false,
		},
	}
	for _, tt := range testCases {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. want %d. got=%d", expected, result.Value)
		return
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. want %t. got=%t", expected, result.Value)
		return
	}
}

func testEval(s string) object.Object {
	l := lexer.New(s)
	p := parser.New(l)
	program := p.ParseProgram()

	return evaluator.Eval(program)
}
