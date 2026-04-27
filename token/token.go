package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Token types
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
	INC      = "++"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"
	LTE = "<="
	GTE = ">="
	LSHIFT = "<<"
	RSHIFT = ">>"
	MOD = "%"

	EQ     = "=="
	NOT_EQ = "!="
	AND    = "&&"
	OR     = "||"

	DOT    = "."
	ARROW  = "->"

	// Delimiters
	COLON     = ":"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

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
	FOR      = "FOR"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	LOAD     = "LOAD"
	STRUCT   = "STRUCT"
	NEW      = "NEW"
)

// Keywords map
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return":   RETURN,
	"print":    PRINT,
	"while":    WHILE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"load":     LOAD,
	"struct":   STRUCT,
	"new":      NEW,
}

// Lookup function
func GetTokenType(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
