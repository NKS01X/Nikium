package lexer

import (
	"Nikium/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `print "Hello World"; 
x = 42i32;
if x == 42 { print "yes"; } else { print "no"; }`

	expectedTokens := []token.Token{
		{Type: "PRINT", Literal: "print"},
		{Type: "STRING", Literal: "Hello World"},
		{Type: "SEMICOLON", Literal: ";"},
		{Type: "IDENT", Literal: "x"},
		{Type: "ASSIGN", Literal: "="},
		{Type: "I32", Literal: "42i32"},
		{Type: "SEMICOLON", Literal: ";"},
		{Type: "IF", Literal: "if"},
		{Type: "IDENT", Literal: "x"},
		{Type: "EQ", Literal: "=="},
		{Type: "INT", Literal: "42"},
		{Type: "LBRACE", Literal: "{"},
		{Type: "PRINT", Literal: "print"},
		{Type: "STRING", Literal: "yes"},
		{Type: "SEMICOLON", Literal: ";"},
		{Type: "RBRACE", Literal: "}"},
		{Type: "ELSE", Literal: "else"},
		{Type: "LBRACE", Literal: "{"},
		{Type: "PRINT", Literal: "print"},
		{Type: "STRING", Literal: "no"},
		{Type: "SEMICOLON", Literal: ";"},
		{Type: "RBRACE", Literal: "}"},
		{Type: "EOF", Literal: ""},
	}

	l := New(input)

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != expected.Type || tok.Literal != expected.Literal {
			t.Fatalf("token[%d] wrong. got=%+v, want=%+v", i, tok, expected)
		}
	}
}
