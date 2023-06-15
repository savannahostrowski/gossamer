package lexer

import (
	"testing"

	"github.com/savannahostrowski/gossamer/token"
)

func TestNextToken(t *testing.T) {
	input := `five = 5
	ten = 10
		
	def add(x, y):
		return x + y
	
	result = add(five, ten)`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.EQUAL, "="},
		{token.NUMBER, "5"},
		{token.IDENT, "ten"},
		{token.EQUAL, "="},
		{token.NUMBER, "10"},
		{token.DEF, "def"},
		{token.IDENT, "add"},
		{token.LPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAR, ")"},
		{token.COLON, ":"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.IDENT, "result"},
		{token.EQUAL, "="},
		{token.IDENT, "add"},
		{token.LPAR, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAR, ")"},
		{token.ENDMARKER, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got= %q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
