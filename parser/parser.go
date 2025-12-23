package parser

import (
	"Nikium/ast"
	"Nikium/lexer"
	"Nikium/token"
	"fmt"
	"strconv"
)

// Parser struct
type Parser struct {
	l       *lexer.Lexer
	curTok  token.Token
	peekTok token.Token
	errors  []string
}

// Operator precedence
var precedences = map[string]int{
	"EQ":     1, // ==
	"NOT_EQ": 1, // !=
	"PLUS":   2,
	"MINUS":  2,
	"MUL":    3,
}

// New creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

// advance tokens
func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

// ParseProgram parses the entire input into a Program AST node
func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.curTok.Type != "EOF" {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}
	return prog
}

// Parse a statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case "IDENT":
		return p.parseLetStatement()
	case "PRINT":
		return p.parsePrintStatement()
	case "IF":
		return p.parseIfStatement()
	default:
		msg := fmt.Sprintf("unknown statement: %s", p.curTok.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
}

// Assignment: x = 42i32;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curTok}
	name := &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	stmt.Name = name

	p.nextToken()
	if p.curTok.Type != "ASSIGN" {
		p.errors = append(p.errors, "expected ASSIGN after identifier")
		return nil
	}

	p.nextToken()
	val := p.parseExpression(0)
	if val == nil {
		return nil
	}
	stmt.Value = val

	if p.peekTok.Type == "SEMICOLON" {
		p.nextToken()
	}
	return stmt
}

// Print statement: print x; or print "hello";
func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{Token: p.curTok}
	p.nextToken()
	val := p.parseExpression(0)
	if val == nil {
		p.errors = append(p.errors, "expected expression after print")
		return nil
	}
	stmt.Value = val

	if p.peekTok.Type == "SEMICOLON" {
		p.nextToken()
	}
	return stmt
}

// If statement: if cond { ... } else { ... }
func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curTok}
	p.nextToken()
	stmt.Condition = p.parseExpression(0)

	if p.curTok.Type != "LBRACE" {
		p.errors = append(p.errors, "expected { after if condition")
		return nil
	}
	stmt.Consequence = p.parseBlockStatement()

	if p.peekTok.Type == "ELSE" {
		p.nextToken() // consume ELSE
		p.nextToken() // move to LBRACE
		stmt.Alternative = p.parseBlockStatement()
	}
	return stmt
}

// Block: { ... }
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curTok}
	block.Statements = []ast.Statement{}

	p.nextToken()
	for p.curTok.Type != "RBRACE" && p.curTok.Type != "EOF" {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

// Parse expressions with precedence
func (p *Parser) parseExpression(precedence int) ast.Expression {
	var left ast.Expression

	switch p.curTok.Type {
	case "IDENT":
		left = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	case "INT":
		v, _ := strconv.ParseInt(p.curTok.Literal, 10, 64)
		left = &ast.IntegerLiteral{Token: p.curTok, Value: v}
	case "I32", "I64":
		numStr := p.curTok.Literal[:len(p.curTok.Literal)-3]
		v, _ := strconv.ParseInt(numStr, 10, 64)
		left = &ast.IntegerLiteral{Token: p.curTok, Value: v}
	case "STRING":
		left = &ast.StringLiteral{Token: p.curTok, Value: p.curTok.Literal}
	case "LPAREN":
		p.nextToken()
		left = p.parseExpression(0)
		if p.curTok.Type != "RPAREN" {
			p.errors = append(p.errors, "expected )")
			return nil
		}
	default:
		p.errors = append(p.errors, "unknown expression: "+p.curTok.Literal)
		return nil
	}

	for p.peekTok.Type != "SEMICOLON" && precedence < p.peekPrecedence() {
		op := p.peekTok
		if _, ok := precedences[op.Type]; !ok {
			return left
		}
		p.nextToken()
		left = p.parseInfixExpression(left)
	}
	return left
}

// Parse infix expression: x + 10
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.BinaryExpression{
		Token:    p.curTok,
		Operator: p.curTok.Literal,
		Left:     left,
	}
	prec := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(prec)
	return exp
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekTok.Type]; ok {
		return p
	}
	return 0
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curTok.Type]; ok {
		return p
	}
	return 0
}

// Errors returns parser errors
func (p *Parser) Errors() []string {
	return p.errors
}
