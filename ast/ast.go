package ast

import (
	"Nikium/token"
	"fmt"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	out := ""
	for _, s := range p.Statements {
		out += s.String() + "\n"
	}
	return out
}

// ---------------- Identifier ----------------
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ---------------- LetStatement ----------------
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Type  string // optional type annotation
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	out := ls.Name.String()
	if ls.Type != "" {
		out += ":" + ls.Type
	}
	out += " = " + ls.Value.String()
	return out
}

// ---------------- PrintStatement ----------------
type PrintStatement struct {
	Token token.Token
	Value Expression
}

func (ps *PrintStatement) statementNode()       {}
func (ps *PrintStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PrintStatement) String() string {
	return "print " + ps.Value.String()
}

// ---------------- Literals ----------------
type IntegerLiteral struct {
	Token token.Token
	Value int64
	Type  string
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return `"` + sl.Value + `"` }

// ---------------- BinaryExpression ----------------
type BinaryExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }
func (be *BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", be.Left.String(), be.Operator, be.Right.String())
}

// ---------------- BlockStatement ----------------
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	out := ""
	for _, s := range bs.Statements {
		out += "    " + s.String() + "\n"
	}
	return out
}

// ---------------- IfStatement ----------------
type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	out := "if " + is.Condition.String() + ":\n" + is.Consequence.String()
	if is.Alternative != nil {
		out += "else:\n" + is.Alternative.String()
	}
	return out
}

// ---------------- WhileStatement ----------------
type WhileStatement struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	out := "while " + ws.Condition.String() + ":\n" + ws.Body.String()
	return out
}
