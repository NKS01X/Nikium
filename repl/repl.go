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
  - with type: my_var:i32 = 10;
  - without type: my_var = 10;

- Data Types:
  - i32, i64, string

- Operators:
  - Arithmetic: +, -, *
  - Relational: ==, !=, <, >

- Control Flow:
  - if-else statements
  - while loops

- Built-in Functions:
  - print: print "Hello";

- Comments:
  - Not supported

- Escape Characters in Strings:
  - \\n, \\t, \\\\, \\"
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := evaluator.NewEnvironment()
	fmt.Fprintln(out, BANNER)
	fmt.Fprintln(out, FEATURES)
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
