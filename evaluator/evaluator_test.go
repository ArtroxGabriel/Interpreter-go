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
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testIntegerObject(t, evaluated, tC.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 > 1", false},
		{"1 < 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}
	for _, tt := range testCases {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "!true",
			expected: false,
		},
		{
			input:    "!false",
			expected: true,
		},
		{
			input:    "!5",
			expected: false,
		},
		{
			input:    "!!true",
			expected: true,
		},
		{
			input:    "!!false",
			expected: false,
		},
		{
			input:    "!!5",
			expected: true,
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
