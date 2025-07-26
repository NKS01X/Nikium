package lexer

type Lexer struct {
	input   string
	currpos int
	currinp int
	ch      byte
}

func new(inp string) *Lexer {
	l := &Lexer{input: inp}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.currinp >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.currinp]
	}
	l.currpos = l.currinp
	l.currinp++
}
