package parser

import (
	"Nikium/ast"
	"Nikium/lexer"
	"Nikium/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[string]int{
	"EQ":       EQUALS,
	"NOT_EQ":   EQUALS,
	"LT":       LESSGREATER,
	"GT":       LESSGREATER,
	"PLUS":     SUM,
	"MINUS":    SUM,
	"SLASH":    PRODUCT,
	"ASTERISK": PRODUCT,
	"LPAREN":   CALL,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[string]prefixParseFn
	infixParseFns  map[string]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[string]prefixParseFn)
	p.registerPrefix("IDENT", p.parseIdentifier)
	p.registerPrefix("INT", p.parseIntegerLiteral)
	p.registerPrefix("STRING", p.parseStringLiteral)
	p.registerPrefix("TRUE", p.parseBoolean)
	p.registerPrefix("FALSE", p.parseBoolean)
	p.registerPrefix("BANG", p.parsePrefixExpression)
	p.registerPrefix("MINUS", p.parsePrefixExpression)
	p.registerPrefix("LPAREN", p.parseGroupedExpression)
	p.registerPrefix("IF", p.parseIfStatement)
	p.registerPrefix("WHILE", p.parseWhileStatement)
	p.registerPrefix("FUNCTION", p.parseFunctionLiteral)

	p.infixParseFns = make(map[string]infixParseFn)
	p.registerInfix("PLUS", p.parseInfixExpression)
	p.registerInfix("MINUS", p.parseInfixExpression)
	p.registerInfix("SLASH", p.parseInfixExpression)
	p.registerInfix("ASTERISK", p.parseInfixExpression)
	p.registerInfix("EQ", p.parseInfixExpression)
	p.registerInfix("NOT_EQ", p.parseInfixExpression)
	p.registerInfix("LT", p.parseInfixExpression)
	p.registerInfix("GT", p.parseInfixExpression)
	p.registerInfix("LPAREN", p.parseCallExpression)

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.curToken.Type != "EOF" {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case "IDENT":
		if p.peekToken.Type == "ASSIGN" || p.peekToken.Type == "COLON" {
			return p.parseLetStatement()
		}
		return p.parseExpressionStatement()
	case "PRINT":
		return p.parsePrintStatement()
	case "RETURN":
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type == "COLON" {
		p.nextToken()
		p.nextToken()
		stmt.Type = p.curToken.Literal
	}

	if !p.expectPeek("ASSIGN") {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs("SEMICOLON") {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs("SEMICOLON") {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekTokenIs("SEMICOLON") {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs("SEMICOLON") {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.errors = append(p.errors, "no prefix parse function for "+p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs("SEMICOLON") && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

/* ---------- expressions ---------- */

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, "invalid integer")
		return nil
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: val}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs("TRUE"),
	}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.BinaryExpression{
		Token:    p.curToken,
		Left:     left,
		Operator: p.curToken.Literal,
	}
	prec := p.curPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(prec)
	return expr
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek("RPAREN") {
		return nil
	}
	return exp
}

/* ---------- control ---------- */

func (p *Parser) parseIfStatement() ast.Expression {
	expr := &ast.IfStatement{Token: p.curToken}
	p.nextToken()
	expr.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek("LBRACE") {
		return nil
	}
	expr.Consequence = p.parseBlockStatement()

	if p.peekTokenIs("ELSE") {
		p.nextToken()
		if !p.expectPeek("LBRACE") {
			return nil
		}
		expr.Alternative = p.parseBlockStatement()
	}
	return expr
}

func (p *Parser) parseWhileStatement() ast.Expression {
	expr := &ast.WhileStatement{Token: p.curToken}
	p.nextToken()
	expr.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek("LBRACE") {
		return nil
	}
	expr.Body = p.parseBlockStatement()
	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs("RBRACE") && !p.curTokenIs("EOF") {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

/* ---------- functions & calls ---------- */

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek("LPAREN") {
		return nil
	}
	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek("LBRACE") {
		return nil
	}
	lit.Body = p.parseBlockStatement()
	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	params := []*ast.Identifier{}

	if p.peekTokenIs("RPAREN") {
		p.nextToken()
		return params
	}

	p.nextToken()
	params = append(params, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})

	for p.peekTokenIs("COMMA") {
		p.nextToken()
		p.nextToken()
		params = append(params, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
	}

	if !p.expectPeek("RPAREN") {
		return nil
	}
	return params
}

func (p *Parser) parseCallExpression(fn ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: fn}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs("RPAREN") {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs("COMMA") {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek("RPAREN") {
		return nil
	}
	return args
}

/* ---------- helpers ---------- */

func (p *Parser) curTokenIs(t string) bool  { return p.curToken.Type == t }
func (p *Parser) peekTokenIs(t string) bool { return p.peekToken.Type == t }

func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors,
		fmt.Sprintf("expected next token %s, got %s", t, p.peekToken.Type))
	return false
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(t string, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfix(t string, fn infixParseFn) {
	p.infixParseFns[t] = fn
}
