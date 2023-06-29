package ast

import (
	"github.com/savannahostrowski/gossamer/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VariableStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "my_var"},
					Value: "my_var",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "another_variable"},
					Value: "another_variable",
				},
			},
		},
	}

	if program.String() != "let my_var = another_variable;" {
		t.Errorf("program.String() wrong. got=%q.", program.String())
	}
}