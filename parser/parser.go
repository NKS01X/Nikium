package parser

import (
	"Nikium/ast"
	"Nikium/lexer"
	"Nikium/token"
	"fmt"
	"strconv"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  token.Token
	peekTok token.Token
	errors  []string
}

// Operator precedence
var precedences = map[string]int{
	"EQ":     1,
	"NOT_EQ": 1,
	"LT":     1,
	"GT":     1,
	"PLUS":   2,
	"MINUS":  2,
	"MUL":    3,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

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

// ---------------- Statements ----------------
func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case "IDENT":
		return p.parseLetStatement()
	case "PRINT":
		return p.parsePrintStatement()
	case "IF":
		return p.parseIfStatement()
	case "WHILE":
		return p.parseWhileStatement()
	case "DEDENT":
		return nil
	default:
		msg := fmt.Sprintf("unknown statement: %s", p.curTok.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
}

// Assignment: a:i32 = 10;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curTok}
	name := &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	stmt.Name = name

	if p.peekTok.Type == "COLON" {
		p.nextToken() // IDENT
		p.nextToken() // COLON

		if p.curTok.Type == "IDENT" && (p.curTok.Literal == "i32" || p.curTok.Literal == "i64" || p.curTok.Literal == "string") {
			stmt.Type = p.curTok.Literal
		} else {
			p.errors = append(p.errors, "expected type after :")
			return nil
		}
	}

	if p.peekTok.Type != "ASSIGN" {
		p.errors = append(p.errors, "expected ASSIGN after identifier/type")
		return nil
	}

	p.nextToken() // consume IDENT or type
	p.nextToken() // consume ASSIGN

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

// Print statement
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

// If statement with indentation
func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curTok}
	p.nextToken()
	stmt.Condition = p.parseExpression(0)

	// Expect COLON after condition
	if p.peekTok.Type != "COLON" {
		p.errors = append(p.errors, "expected COLON after if condition")
		return nil
	}
	p.nextToken() // consume COLON

	// Expect INDENT for block
	if p.peekTok.Type != "INDENT" {
		p.errors = append(p.errors, "expected INDENT after if condition")
		return nil
	}
	p.nextToken() // consume INDENT
	stmt.Consequence = p.parseBlockStatement()

	// Optional else
	if p.peekTok.Type == "ELSE" {
		p.nextToken() // consume ELSE
		// colon after else is optional
		if p.peekTok.Type == "COLON" {
			p.nextToken()
		}

		if p.peekTok.Type != "INDENT" {
			p.errors = append(p.errors, "expected INDENT after else")
			return nil
		}
		p.nextToken()
		stmt.Alternative = p.parseBlockStatement()
	}
	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curTok}
	p.nextToken()
	stmt.Condition = p.parseExpression(0)

	// Expect COLON after condition
	if p.peekTok.Type != "COLON" {
		p.errors = append(p.errors, "expected COLON after while condition")
		return nil
	}
	p.nextToken() // consume COLON

	// Expect INDENT for block
	if p.peekTok.Type != "INDENT" {
		p.errors = append(p.errors, "expected INDENT after while condition")
		return nil
	}
	p.nextToken() // consume INDENT
	stmt.Body = p.parseBlockStatement()

	return stmt
}

// Block: INDENT ... DEDENT
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{}
	block.Statements = []ast.Statement{}

	for p.curTok.Type != "DEDENT" && p.curTok.Type != "EOF" {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

// ---------------- Expressions ----------------
func (p *Parser) parseExpression(precedence int) ast.Expression {
	var left ast.Expression

	switch p.curTok.Type {
	case "IDENT":
		left = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	case "INT":
		v, _ := strconv.ParseInt(p.curTok.Literal, 10, 64)
		left = &ast.IntegerLiteral{Token: p.curTok, Value: v, Type: "i64"}
	case "I32":
		numStr := p.curTok.Literal[:len(p.curTok.Literal)-3]
		v, _ := strconv.ParseInt(numStr, 10, 32)
		left = &ast.IntegerLiteral{Token: p.curTok, Value: v, Type: "i32"}
	case "I64":
		numStr := p.curTok.Literal[:len(p.curTok.Literal)-3]
		v, _ := strconv.ParseInt(numStr, 10, 64)
		left = &ast.IntegerLiteral{Token: p.curTok, Value: v, Type: "i64"}
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

func (p *Parser) Errors() []string {
	return p.errors
}
