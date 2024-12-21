package ast_test

import (
	"testing"

	"github.com/ArtroxGabriel/interpreter/ast"
	"github.com/ArtroxGabriel/interpreter/token"
)

func TestString(t *testing.T) {
	programs := []ast.Program{
		{
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.Token{Type: token.LET, Literal: "let"},
					Name: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "myVar"},
						Value: "myVar",
					},
					Value: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
						Value: "anotherVar",
					},
				},
			},
		},
		{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Token: token.Token{Type: token.RETURN, Literal: "return"},
					ReturnValue: &ast.Identifier{
						Token: token.Token{Type: token.INT, Literal: "10"},
						Value: "10",
					},
				},
			},
		},
	}

	tests := []struct {
		expectedString string
	}{
		{"let myVar = anotherVar;"},
		{"return 10;"},
	}

	for i, tt := range tests {
		program := programs[i]
		if program.String() != tt.expectedString {
			t.Errorf(
				"program.String does not %s. got=%s",
				tt.expectedString,
				program.String(),
			)
		}
	}
}
