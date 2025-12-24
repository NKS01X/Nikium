package lexer

import (
	"Nikium/token"
)

type Lexer struct {
	input       string
	currpos     int
	currinp     int
	ch          byte
	indentStack []int
	atLineStart bool
	emitDEDENT  []token.Token
}

func New(inp string) *Lexer {
	l := &Lexer{
		input:       inp,
		indentStack: []int{0},
		atLineStart: true,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.currinp >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.currinp]
	}
	l.currpos = l.currinp
	l.currinp++
}

func (l *Lexer) PeekChar() byte {
	if l.currinp >= len(l.input) {
		return 0
	}
	return l.input[l.currinp]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) readIndent() int {
	count := 0
	for l.ch == ' ' {
		count++
		l.readChar()
	}
	return count
}

func (l *Lexer) ReadIdentifier() string {
	start := l.currpos
	for token.IsLetter(l.ch) || token.IsDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.currpos]
}

func (l *Lexer) ReadInteger() string {
	start := l.currpos
	for token.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.currpos]
}

func (l *Lexer) ReadString() string {
	l.readChar() // consume initial "
	var out string
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				out += "\n"
			case 't':
				out += "\t"
			case '\\':
				out += "\\"
			case '"':
				out += `"`
			default:
				// Or handle error, for now just append the char
				out += string(l.ch)
			}
		} else {
			out += string(l.ch)
		}
		l.readChar()
	}
	l.readChar() // consume final "
	return out
}

func (l *Lexer) NextToken() token.Token {
	if len(l.emitDEDENT) > 0 {
		tok := l.emitDEDENT[0]
		l.emitDEDENT = l.emitDEDENT[1:]
		return tok
	}

	if l.atLineStart {
		l.atLineStart = false
		indent := l.readIndent()
		lastIndent := l.indentStack[len(l.indentStack)-1]
		if indent > lastIndent {
			l.indentStack = append(l.indentStack, indent)
			return token.Token{Type: "INDENT", Literal: ""}
		} else if indent < lastIndent {
			for indent < l.indentStack[len(l.indentStack)-1] {
				l.indentStack = l.indentStack[:len(l.indentStack)-1]
				l.emitDEDENT = append(l.emitDEDENT, token.Token{Type: "DEDENT", Literal: ""})
			}
			if len(l.emitDEDENT) > 0 {
				tok := l.emitDEDENT[0]
				l.emitDEDENT = l.emitDEDENT[1:]
				return tok
			}
		}
	}

	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.readChar()
			l.atLineStart = true
			return l.NextToken()
		}
		l.readChar()
	}

	// Check for single-character tokens first
	switch l.ch {
	case '=', '!', '+', '-', '*', ';', '(', ')', '{', '}', ':', '<', '>':
		if l.ch == '=' && l.PeekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tokType := token.GetTokenType(literal)
			l.readChar()
			return token.Token{Type: tokType, Literal: literal}
		}
		if l.ch == '!' && l.PeekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tokType := token.GetTokenType(literal)
			l.readChar()
			return token.Token{Type: tokType, Literal: literal}
		}
		tok := token.Token{Type: token.GetTokenType(string(l.ch)), Literal: string(l.ch)}
		l.readChar()
		return tok
	case 0:
		if len(l.indentStack) > 1 {
			l.indentStack = l.indentStack[:len(l.indentStack)-1]
			return token.Token{Type: "DEDENT", Literal: ""}
		}
		return token.Token{Type: "EOF", Literal: ""}
	case '"':
		str := l.ReadString()
		return token.Token{Type: "STRING", Literal: str}
	}

	if token.IsLetter(l.ch) || l.ch == '_' {
		ident := l.ReadIdentifier()
		return token.Token{Type: token.GetTokenType(ident), Literal: ident}
	}

	if token.IsDigit(l.ch) {
		num := l.ReadInteger()
		return token.Token{Type: token.GetTokenType(num), Literal: num}
	}

	ch := l.ch
	l.readChar()
	return token.Token{Type: token.GetTokenType(string(ch)), Literal: string(ch)}
}
