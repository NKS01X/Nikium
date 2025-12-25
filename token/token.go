package token

type TokenType string

type Token struct {
	Type    string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	PRINT    = "PRINT"
	WHILE    = "WHILE"
)

var keywords = map[string]string{
	"fn":     "FUNCTION",
	"let":    "IDENT",
	"true":   "TRUE",
	"false":  "FALSE",
	"if":     "IF",
	"else":   "ELSE",
	"return": "RETURN",
	"print":  "PRINT",
	"while":  "WHILE",
}

func GetTokenType(tok string) string {
	if t, ok := keywords[tok]; ok {
		return t
	}
	return "IDENT"
}
