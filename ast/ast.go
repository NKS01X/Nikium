package ast

import (
	"Nikium/token"
	"fmt"
)

// Node represents any node in the AST. Every node must be able to return
// its token literal and a string representation for debugging or printing.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a type of Node that represents statements in the language.
// Examples: variable assignment, print statements, if statements.
type Statement interface {
	Node
	statementNode()
}

// Expression is a type of Node that represents expressions.
// Examples: identifiers, integer literals, binary expressions.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST. It contains a list of statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal of the first token in the program.
// Used mostly for debugging.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a string representation of the entire program.
func (p *Program) String() string {
	out := ""
	for _, s := range p.Statements {
		out += s.String() + "\n"
	}
	return out
}

// -------------------- Basic AST Nodes --------------------

// Identifier represents variable names or function names.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the actual name of the identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// LetStatement represents a variable assignment like `x = 42`.
type LetStatement struct {
	Token token.Token // the '=' token
	Name  *Identifier // the variable being assigned
	Value Expression  // the expression being assigned to the variable
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	return ls.Name.String() + " = " + ls.Value.String()
}

// PrintStatement represents a print statement like `print x` or `print "hello"`.
type PrintStatement struct {
	Token token.Token // the 'print' token
	Value Expression  // the expression to print
}

func (ps *PrintStatement) statementNode()       {}
func (ps *PrintStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PrintStatement) String() string {
	return "print " + ps.Value.String()
}

// IntegerLiteral represents integer values, including i32 and i64 suffixes.
type IntegerLiteral struct {
	Token token.Token // the token representing the integer
	Value int64       // the numeric value
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// StringLiteral represents string values like `"Hello World"`.
type StringLiteral struct {
	Token token.Token // the '"' token
	Value string      // the string content
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return `"` + sl.Value + `"` }

// -------------------- Expressions --------------------

// BinaryExpression represents an operation between two expressions like `x + y`.
type BinaryExpression struct {
	Token    token.Token // operator token, e.g., '+'
	Left     Expression  // left-hand side expression
	Operator string      // operator literal
	Right    Expression  // right-hand side expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }
func (be *BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", be.Left.String(), be.Operator, be.Right.String())
}

// -------------------- Control Flow --------------------

// BlockStatement represents a group of statements inside `{ ... }`.
type BlockStatement struct {
	Token      token.Token // '{' token
	Statements []Statement // list of statements inside the block
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	out := ""
	for _, s := range bs.Statements {
		out += s.String() + "\n"
	}
	return out
}

// IfStatement represents an `if` or `if-else` statement.
type IfStatement struct {
	Token       token.Token     // 'if' token
	Condition   Expression      // condition expression
	Consequence *BlockStatement // block executed if condition is true
	Alternative *BlockStatement // block executed if condition is false (optional)
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	out := "if " + is.Condition.String() + " {\n" + is.Consequence.String() + "}"
	if is.Alternative != nil {
		out += " else {\n" + is.Alternative.String() + "}"
	}
	return out
}
