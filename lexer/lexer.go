package lexer

import (
	"Nikium/token"
	"fmt"
)

// Lexer breaks input into tokens
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
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) ReadIdentifier() string {
	idx := l.currpos
	for token.IsLetter(l.ch) || token.IsDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[idx:l.currpos]
}

func (l *Lexer) ReadInteger() string {
	idx := l.currpos
	for token.IsDigit(l.ch) {
		l.readChar()
	}

	// handle suffixes i32, i64
	if l.ch == 'i' {
		l.readChar()
		for token.IsLetter(l.ch) || token.IsDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[idx:l.currpos]
}

func (l *Lexer) ReadString() string {
	l.readChar() // skip opening quote
	idx := l.currpos
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[idx:l.currpos]
	l.readChar() // skip closing quote
	return str
}

// printToken prints token in language-style format
func (l *Lexer) printToken(tok token.Token) {
	fmt.Printf("<%s : %s>\n", tok.Type, tok.Literal)
}

// NextToken returns the next token
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
			tok := token.Token{Type: tokType, Literal: literal}
			l.printToken(tok)
			return tok
		}
	}

	// String literal
	if l.ch == '"' {
		str := l.ReadString()
		tok := token.Token{Type: "STRING", Literal: str}
		l.printToken(tok)
		return tok
	}

	// Identifier / Keyword
	if token.IsLetter(l.ch) || l.ch == '_' {
		ident := l.ReadIdentifier()
		tokType := token.GetTokenType(ident)
		tok := token.Token{Type: tokType, Literal: ident}
		l.printToken(tok)
		return tok
	}

	// EOF
	if l.ch == 0 {
		tok := token.Token{Type: "EOF", Literal: ""}
		l.printToken(tok)
		return tok
	}

	// Number
	if token.IsDigit(l.ch) {
		num := l.ReadInteger()
		tokType := token.GetTokenType(num)
		tok := token.Token{Type: tokType, Literal: num}
		l.printToken(tok)
		return tok
	}

	// Single-character symbols
	ch := l.ch
	l.readChar()
	literal := string(ch)
	tokType := token.GetTokenType(literal)
	tok := token.Token{Type: tokType, Literal: literal}
	l.printToken(tok)
	return tok
}
