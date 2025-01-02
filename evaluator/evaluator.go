package evaluator

import (
	"github.com/ArtroxGabriel/interpreter/ast"
	"github.com/ArtroxGabriel/interpreter/object"
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

	case *ast.BlockStatement:
	case *ast.Boolean:
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
