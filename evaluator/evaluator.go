package evaluator

import (
	"github.com/ArtroxGabriel/interpreter/ast"
	"github.com/ArtroxGabriel/interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// Statemennts
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.BlockStatement:
	case *ast.CallExpression:
	case *ast.FunctionalLiteral:
	case *ast.Identifier:
	case *ast.IfExpression:
	case *ast.InfixExpression:
	case *ast.LetStatement:
	case *ast.PrefixExpression:
	case *ast.ReturnStatement:
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
