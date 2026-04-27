package main

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"Nikium/repl"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	env := evaluator.NewEnvironment()

	if len(os.Args) > 1 {
		filePath := os.Args[1]

		content, err := ioutil.ReadFile(filePath)
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

		evaluator.Eval(program, env)
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
