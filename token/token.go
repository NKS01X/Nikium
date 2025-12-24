package token

import "strings"

type Token struct {
	Type    string
	Literal string
}

// TokenLiterals maps keywords, symbols, operators
var TokenLiterals = map[string]string{
	"I32":   "I32",
	"I64":   "I64",
	"if":    "IF",
	"else":  "ELSE",
	"true":  "TRUE",
	"false": "FALSE",
	"print": "PRINT",
	"while": "WHILE",
	"=":     "ASSIGN",
	"+":     "PLUS",
	"-":     "MINUS",
	"*":     "MUL",
	"==":    "EQ",
	"!=":    "NOT_EQ",
	"<":     "LT",
	">":     "GT",
	";":     "SEMICOLON",
	"(":     "LPAREN",
	")":     "RPAREN",
	"{":     "LBRACE",
	"}":     "RBRACE",
	":":     "COLON",
}

// GetTokenType returns the token type for a literal
func GetTokenType(literal string) string {
	if tokType, ok := TokenLiterals[literal]; ok {
		return tokType
	}

	if isIdentifier(literal) {
		return "IDENT"
	}

	if strings.HasSuffix(literal, "i32") {
		numPart := literal[:len(literal)-3]
		if isNumber(numPart) {
			return "I32"
		}
		return "ILLEGAL"
	}

	if strings.HasSuffix(literal, "i64") {
		numPart := literal[:len(literal)-3]
		if isNumber(numPart) {
			return "I64"
		}
		return "ILLEGAL"
	}

	if isNumber(literal) {
		return "INT"
	}

	return "ILLEGAL"
}

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isIdentifier(s string) bool {
	if len(s) == 0 || IsDigit(s[0]) {
		return false
	}
	if _, isKeyword := TokenLiterals[s]; isKeyword {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !IsLetter(s[i]) && !IsDigit(s[i]) && s[i] != '_' {
			return false
		}
	}
	return true
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func IsLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
}
