package main

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"Nikium/repl"
	"fmt"
	"os"
)

func main() {
	env := evaluator.NewEnvironment()

	if len(os.Args) > 1 {
		filePath := os.Args[1]

		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
			os.Exit(1)
		}

		l := lexer.New(string(content))
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Fprintln(os.Stderr, "Parser error:", err)
			}
			os.Exit(1)
		}

		result := evaluator.Eval(program, env)
		if result != nil && result.Type() == evaluator.ERROR_OBJ {
			fmt.Fprintln(os.Stderr, result.Inspect())
			os.Exit(1)
		}
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
