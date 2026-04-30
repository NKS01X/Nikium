package repl

import (
	"Nikium/evaluator"
	"Nikium/lexer"
	"Nikium/parser"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
	ColorDim    = "\033[2m"
)

const PROMPT = ColorCyan + "nikium> " + ColorReset

const BANNER = ColorCyan + ColorBold + `
   ╔══════════════════════════════════════════════════╗
   ║  _   _ _ _    _                                  ║
   ║ | \ | (_) |  (_)_   _ _ __ ___                   ║
   ║ |  \| | | |/ /| | | | | '_ ` + "`" + ` _ \                  ║
   ║ | |\  | |   < | | |_| | | | | | |                 ║
   ║ |_| \_|_|_|\_\|_|\__,_|_| |_| |_|                 ║
   ║                                                  ║
   ║      Advanced Agentic Coding Environment         ║
   ╚══════════════════════════════════════════════════╝
` + ColorReset

const HELP = ColorGreen + `
  Available Commands:
   help     - Show this help message
   clear    - Clear the terminal screen
   features - List language features
   exit     - Exit Nikium REPL
` + ColorReset

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

	// Register REPL commands as builtins too, so they work in expressions
	env.Set("help", &evaluator.Function{
		Native: func(args []evaluator.Object) evaluator.Object {
			fmt.Fprint(out, HELP)
			return evaluator.NULL
		},
	})
	env.Set("clear", &evaluator.Function{
		Native: func(args []evaluator.Object) evaluator.Object {
			fmt.Fprint(out, "\033[H\033[2J")
			return evaluator.NULL
		},
	})
	env.Set("exit", &evaluator.Function{
		Native: func(args []evaluator.Object) evaluator.Object {
			fmt.Fprintln(out, ColorYellow+"  Goodbye! 👋"+ColorReset)
			os.Exit(0)
			return evaluator.NULL
		},
	})

	fmt.Fprint(out, BANNER)
	fmt.Fprintln(out, ColorDim+"  Type 'help' for commands or 'exit' to quit.\n"+ColorReset)

	var lines []string
	bracketCount := 0

	for {
		if bracketCount == 0 {
			fmt.Fprint(out, ColorBlue+ColorBold+" nikium "+ColorReset+ColorCyan+" "+ColorReset)
		} else {
			fmt.Fprint(out, ColorCyan+strings.Repeat("  ", bracketCount)+".. "+ColorReset)
		}

		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if trimmed == "" && bracketCount == 0 {
			continue
		}

		// Handle REPL Commands (only if not in multiline)
		if bracketCount == 0 {
			switch strings.ToLower(trimmed) {
			case "exit", "quit", ":q":
				fmt.Fprintln(out, ColorYellow+"  Goodbye! 👋"+ColorReset)
				return
			case "help", "?":
				fmt.Fprint(out, HELP)
				continue
			case "clear", "cls":
				fmt.Fprint(out, "\033[H\033[2J")
				continue
			case "features":
				fmt.Fprintln(out, ColorYellow+FEATURES+ColorReset)
				continue
			}
		}

		// Multiline tracking
		bracketCount += strings.Count(line, "{") - strings.Count(line, "}")
		bracketCount += strings.Count(line, "(") - strings.Count(line, ")")
		bracketCount += strings.Count(line, "[") - strings.Count(line, "]")
		if bracketCount < 0 {
			bracketCount = 0
		}

		lines = append(lines, line)

		if bracketCount > 0 {
			continue
		}

		fullInput := strings.Join(lines, "\n")
		lines = nil // reset

		l := lexer.New(fullInput)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			fmt.Fprint(out, ColorRed)
			for _, err := range p.Errors() {
				fmt.Fprintln(out, "  ✘ Parser error:", err)
			}
			fmt.Fprint(out, ColorReset)
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			color := ColorPurple
			switch evaluated.Type() {
			case evaluator.INTEGER_OBJ:
				color = ColorYellow
			case evaluator.BOOLEAN_OBJ:
				color = ColorCyan
			case evaluator.STRING_OBJ:
				color = ColorGreen
			case evaluator.ERROR_OBJ:
				color = ColorRed
			}
			fmt.Fprint(out, color+ColorBold+"   "+ColorReset+color+evaluated.Inspect()+"\n"+ColorReset)
		}
	}
}
