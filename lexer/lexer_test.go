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
	
	result = add(five, ten)

	10 == 10
	10 != 9

	/=
	/
	//
	//=
	*
	**=
	**
	*=
	<
	<<
	<<=
	<=
	>
	>>
	>>=
	>=
	@=
	@
	:
	:=
	.
	...
	->
	-
	`

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
		{token.NUMBER, "10"},
		{token.EQEQUAL, "=="},
		{token.NUMBER, "10"},
		{token.NUMBER, "10"},
		{token.NOTEQUAL, "!="},
		{token.NUMBER, "9"},
		{token.SLASHEQUAL, "/="},
		{token.SLASH, "/"},
		{token.DOUBLESLASH, "//"},
		{token.DOUBLESLASHEQUAL, "//="},
		{token.STAR, "*"},
		{token.DOUBLESTAREQUAL, "**="},
		{token.DOUBLESTAR, "**"},
		{token.STAREQUAL, "*="},
		{token.LESS, "<"},
		{token.LEFTSHIFT, "<<"},
		{token.LEFTSHIFTEQUAL, "<<="},
		{token.LESSEQUAL, "<="},
		{token.GREATER, ">"},
		{token.RIGHTSHIFT, ">>"},
		{token.RIGHTSHIFTEQUAL, ">>="},
		{token.GREATEREQUAL, ">="},
		{token.ATEQUAL, "@="},
		{token.AT, "@"},
		{token.COLON, ":"},
		{token.COLONEQUAL, ":="},
		{token.DOT, "."},
		{token.ELLIPSIS, "..."},
		{token.RARROW, "->"},
		{token.MINUS, "-"},
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
