package repl

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>> "

const BANNER = `
NN    N  II  K   K  H   H  II  L
N N   N  II  K  K   H   H  II  L
N  N  N  II  KK     HHHHH  II  L
N   N N  II  K  K   H   H  II  L
N    NN  II  K   K  H   H  II  LLLLL
`

const FEATURES = `
Nikium Language Features:

- Variable Declaration:
  - with type: let my_var:i32 = 10;
  - without type: let my_var = 10;

- Data Types:
  - int, string, boolean, array, hash, function

- Operators:
  - Arithmetic: +, -, *, /, %, <<, >>
  - Relational: ==, !=, <, >, <=, >=
  - Logical: &&, ||

- Control Flow:
  - if-else statements
  - while loops with break/continue

- Built-in Functions:
  - print: print "Hello";

- Comments:
  - Supported: // this is a comment

- Escape Characters in Strings:
  - \n, \t, \\, \"
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := evaluator.NewEnvironment()
	fmt.Fprint(out, BANNER)
	fmt.Fprint(out, FEATURES)
	fmt.Fprintln(out, "Welcome to Nikium REPL! Type code and press Enter.")

	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" || line == "quit" {
			fmt.Fprintln(out, "Goodbye!")
			break
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Fprintln(out, "Parser error:", err)
			}
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprintln(out, evaluated.Inspect())
		}
	}
}
