package token

import "strings"

//here we have just made a map that will map type_of_token to literal
//we have used map which will give the result in O(1)

// Token represents a lexical token
type Token struct {
	Type    string
	Literal string
}

// TokenLiterals maps known literals to token types
//function and all later for GOD'S sake
var TokenLiterals = map[string]string{
	"I32":   "I32",
	"I64":   "I64",
	"if":    "IF",
	"else":  "ELSE",
	"true":  "TRUE",
	"false": "FALSE",
	"=":     "ASSIGN",
	"+":     "PLUS",
	"-":     "MINUS",
	"*":     "MUL",
	"==":    "EQ",
	"!=":    "NOT_EQ",
}

// GetTokenType returns the token type for a given literal

func GetTokenType(literal string) string {
	if tokType, ok := TokenLiterals[literal]; ok {
		return tokType
	}
	if isIdentifier(literal) {
		return "IDENT"
	}

	// Handle integer literals with suffix
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

	// Plain number (no suffix)
	if isNumber(literal) {
		return "INT"
	}

	return "ILLEGAL"
}

// isNumber checks if a string contains only digits
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

// isDigit checks if a byte is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter checks if a byte is a valid letter or underscore
func IsLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
}
