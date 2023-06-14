package token

type TokenType string

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	IDENT    = "IDENT"
	INT      = "INT"
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	COMMA    = ","
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	COLON    = ":"
	DEF      = "DEF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	ELIF     = "ELIF"
	RETURN   = "RETURN"
	EQ       = "=="
	NOT_EQ   = "!="
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def":    DEF,
	"True":   TRUE,
	"False":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"elif":   ELIF,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
