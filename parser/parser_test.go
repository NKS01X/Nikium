package parser

import (
	"Nikium/ast"
	"Nikium/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `a:i32 = 10;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.LetStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "a" {
		t.Errorf("stmt.Name.Value not 'a'. got=%s", stmt.Name.Value)
	}

	if stmt.Type != "i32" {
		t.Errorf("stmt.Type not 'i32'. got=%s", stmt.Type)
	}

// TODO: Test the value of the expression
}

func TestIfStatementParsing(t *testing.T) {
	input := `
if x > y:
    print "hello";
else:
    print "world";`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IfStatement. got=%T", program.Statements[0])
	}

	if stmt.Condition.String() != "(x > y)" {
		t.Errorf("stmt.Condition.String() wrong. got=%q", stmt.Condition.String())
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(stmt.Consequence.Statements))
	}

	consequence, ok := stmt.Consequence.Statements[0].(*ast.PrintStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.PrintStatement. got=%T", stmt.Consequence.Statements[0])
	}

	if consequence.Value.String() != `"hello"` {
		t.Errorf("consequence.Value.String() wrong. got=%q", consequence.Value.String())
	}

	if stmt.Alternative == nil {
		t.Fatalf("stmt.Alternative is nil")
	}

	if len(stmt.Alternative.Statements) != 1 {
		t.Errorf("alternative is not 1 statements. got=%d\n", len(stmt.Alternative.Statements))
	}

	alternative, ok := stmt.Alternative.Statements[0].(*ast.PrintStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.PrintStatement. got=%T", stmt.Alternative.Statements[0])
	}

	if alternative.Value.String() != `"world"` {
		t.Errorf("alternative.Value.String() wrong. got=%q", alternative.Value.String())
	}
}

func TestWhileStatementParsing(t *testing.T) {
	input := `
while x < 10:
    print x;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.WhileStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.WhileStatement. got=%T", program.Statements[0])
	}

	if stmt.Condition.String() != "(x < 10)" {
		t.Errorf("stmt.Condition.String() wrong. got=%q", stmt.Condition.String())
	}

	if len(stmt.Body.Statements) != 1 {
		t.Errorf("body is not 1 statements. got=%d\n", len(stmt.Body.Statements))
	}

	body, ok := stmt.Body.Statements[0].(*ast.PrintStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.PrintStatement. got=%T", stmt.Body.Statements[0])
	}

	if body.Value.String() != `x` {
		t.Errorf("body.Value.String() wrong. got=%q", body.Value.String())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
