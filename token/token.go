package token

//here we have just made a map that will map type_of_token to literal
//we have used map which will give the result in O(1)

type Token struct {
	Type    string
	Literal string
}

// TokenLiterals maps known literals to token types
//tokens like for and while or someother will be added later for GOD'S sake
var TokenLiterals = map[string]string{
	"let":   "LET",
	"var":   "VAR",
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

// Returns the token type 
func GetTokenType(literal string) string {
	if tokType, ok := TokenLiterals[literal]; ok {
		return tokType
	}
	if isIdentifier(literal) {
		return "IDENT"
	}
	if isNumber(literal) {
		return "INT"
	}
	return "ILLEGAL"
}
 
func isNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}

//check for identiffier
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
