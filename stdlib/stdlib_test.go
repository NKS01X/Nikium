package stdlib

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"testing"
)

func TestAbs(t *testing.T) {
	input := `abs(-10)`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := evaluator.NewEnvironment()
	Register(env)

	evaluated := evaluator.Eval(program, env)
	if evaluated.Inspect() != "10" {
		t.Errorf("expected 10, got %s", evaluated.Inspect())
	}
}