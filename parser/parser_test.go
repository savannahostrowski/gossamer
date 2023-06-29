package parser

import (
	"testing"

	"github.com/savannahostrowski/gossamer/ast"
	"github.com/savannahostrowski/gossamer/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
x = 5
y = 10
foobar = 838383
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
		len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !TestVariableStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestVariableStatements(t *testing.T, s ast.Statement, name string) bool {
	varStmt, ok := s.(*ast.VariableStatement)
	if !ok {
		t.Errorf("s not *ast.VariableStatement. got=%T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s'. got=%s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name.TokenLiteral() not '%s'. got=%s", name, varStmt.Name.TokenLiteral())
		return false
	}
	return true
}