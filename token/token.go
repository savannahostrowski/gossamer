package token

type TokenType string

const (
	ENDMARKER = "ENDMARKER"
	// NAME      = "NAME"
	NUMBER = "NUMBER"
	// STRING    = "STRING"
	// NEWLINE   = "NEWLINE"
	// INDENT    = "INDENT"
	// DEDENT    = "DEDENT"
	IDENT = "IDENT"

	LPAR             = "("
	RPAR             = ")"
	LSQB             = "["
	RSQB             = "]"
	COLON            = ":"
	COMMA            = ","
	SEMI             = ";"
	PLUS             = "+"
	MINUS            = "-"
	STAR             = "*"
	SLASH            = "/"
	VBAR             = "|"
	AMPER            = "&"
	LESS             = "<"
	GREATER          = ">"
	EQUAL            = "="
	DOT              = "."
	PERCENT          = "%"
	LBRACE           = "{"
	RBRACE           = "}"
	EXCLAMATION      = "!"
	EQEQUAL          = "=="
	NOTEQUAL         = "!="
	LESSEQUAL        = "<="
	GREATEREQUAL     = ">="
	TILDE            = "~"
	CIRCUMFLEX       = "^"
	LEFTSHIFT        = "<<"
	RIGHTSHIFT       = ">>"
	DOUBLESTAR       = "**"
	PLUSEQUAL        = "+="
	MINUSEQUAL       = "-="
	STAREQUAL        = "*="
	SLASHEQUAL       = "/="
	PERCENTEQUAL     = "%="
	AMPEREQUAL       = "&="
	VBAREQUAL        = "|="
	CIRCUMFLEXEQUAL  = "^="
	LEFTSHIFTEQUAL   = "<<="
	RIGHTSHIFTEQUAL  = ">>="
	DOUBLESTAREQUAL  = "**="
	DOUBLESLASH      = "//"
	DOUBLESLASHEQUAL = "//="
	AT               = "@"
	ATEQUAL          = "@="
	RARROW           = "->"
	ELLIPSIS         = "..."
	COLONEQUAL       = ":="

	// OP             = "OP"
	// AWAIT          = "AWAIT"
	// ASYNC          = "ASYNC"
	// TYPE_IGNORE    = "TYPE_IGNORE"
	// TYPE_COMMENT   = "TYPE_COMMENT"
	// SOFT_KEYWORD   = "SOFT_KEYWORD"
	// FSTRING_START  = "FSTRING_START"
	// FSTRING_MIDDLE = "FSTRING_MIDDLE"
	// FSTRING_END    = "FSTRING_END"
	// COMMENT        = "COMMENT"
	// NL             = "NL"
	ERRORTOKEN = "ERRORTOKEN"
	// ENCODING       = "ENCODING"

	DEF    = "DEF"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	ELIF   = "ELIF"
	RETURN = "RETURN"
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
