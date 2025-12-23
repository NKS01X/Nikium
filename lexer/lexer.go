package lexer

import (
	"Nikium/token"
)

// lexer is actually a type of scanner that breaks every line of the code into a lexer type struct
type Lexer struct {
	input   string
	currpos int
	currinp int
	ch      byte
}

func New(inp string) *Lexer {
	l := &Lexer{input: inp}
	l.readChar()
	return l
}

func (l *Lexer) readChar() { // reads
	if l.currinp >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.currinp]
	}
	l.currpos = l.currinp
	l.currinp++
}

// now we break every type to token using GetTokenType function

func (l *Lexer) ReadIdentifier() string { // identifier read karega
	idx := l.currpos
	for token.IsDigit(l.ch) || token.IsLetter(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[idx:l.currpos]
}

func (l *Lexer) ReadInteger() string {
	idx := l.currpos
	for token.IsDigit(l.ch) {
		l.readChar()
	}

	// Handle suffixes like i32, i64
	if l.ch == 'i' {
		// start := l.currpos
		l.readChar()
		for token.IsLetter(l.ch) || token.IsDigit(l.ch) {
			l.readChar()
		}
		return l.input[idx:l.currpos]
	}

	return l.input[idx:l.currpos]
}

func (l *Lexer) skipWhitespace() {
	//i := 0
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
		//	i++
	}
	//fmt.Printf("cnt: %d\n", i)
}

func (l *Lexer) PeekChar() byte {
	if l.currinp >= len(l.input) {
		return 0
	}
	return l.input[l.currinp]
}

// THINGS TO DO : 27:07:25
// NEXTTOKEN FUNC IMPLEMENTATION
// CHECKING IF THE TOKEN AND LEXER THINGS IS WORKING OR NOT

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	// Multi-character operators: ==, !=
	if l.ch == '=' || l.ch == '!' {
		if l.PeekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tokType := token.GetTokenType(literal)
			l.readChar()
			return token.Token{Type: tokType, Literal: literal}
		}
	}

	if token.IsLetter(l.ch) || l.ch == '_' {
		ident := l.ReadIdentifier()
		tokType := token.GetTokenType(ident)
		return token.Token{Type: tokType, Literal: ident}
	}

	if l.ch == 0 {
		return token.Token{Type: "EOF", Literal: ""}
	}

	if token.IsDigit(l.ch) {
		num := l.ReadInteger()
		tokType := token.GetTokenType(num)
		return token.Token{Type: tokType, Literal: num}
	}

	// Single-character symbols
	ch := l.ch
	l.readChar()
	literal := string(ch)
	tokType := token.GetTokenType(literal)

	return token.Token{Type: tokType, Literal: literal}
}
