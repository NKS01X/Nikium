package token

import "strings"

// Token represents a lexical token
type Token struct {
	Type    string
	Literal string
}

// TokenLiterals maps known literals/keywords to token types
var TokenLiterals = map[string]string{
	"I32":   "I32",
	"I64":   "I64",
	"if":    "IF",
	"else":  "ELSE",
	"true":  "TRUE",
	"false": "FALSE",
	"print": "PRINT",
	"=":     "ASSIGN",
	"+":     "PLUS",
	"-":     "MINUS",
	"*":     "MUL",
	"==":    "EQ",
	"!=":    "NOT_EQ",
	";":     "SEMICOLON",
	"(":     "LPAREN",
	")":     "RPAREN",
	"{":     "LBRACE",
	"}":     "RBRACE",
}

// GetTokenType returns the token type for a given literal
func GetTokenType(literal string) string {
	if tokType, ok := TokenLiterals[literal]; ok {
		return tokType
	}

	// Identifier
	if isIdentifier(literal) {
		return "IDENT"
	}

	// Integer with suffix
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

	// Plain number
	if isNumber(literal) {
		return "INT"
	}

	// Unknown token
	return "ILLEGAL"
}

// isNumber checks if a string is numeric
func isNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}

// isIdentifier checks if a string is a valid identifier
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

// IsDigit checks if a byte is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// IsLetter checks if a byte is a letter or underscore
func IsLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
}
