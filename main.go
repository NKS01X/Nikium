package main

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"Nikium/repl"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	env := evaluator.NewEnvironment()

	// Load all files in stdlib folder
	stdlibFiles, err := ioutil.ReadDir("stdlib")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdlib folder: %s\n", err)
		os.Exit(1)
	}

	for _, f := range stdlibFiles {
		if f.IsDir() {
			continue
		}

		path := filepath.Join("stdlib", f.Name())
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdlib file %s: %s\n", path, err)
			os.Exit(1)
		}

		l := lexer.New(string(content))
		p := parser.New(l)
		program := p.ParseProgram()
		evaluator.Eval(program, env)
	}

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
