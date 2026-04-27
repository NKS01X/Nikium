import re

with open('/media/nikhil/Windows/Nikium/lexer/lexer.go', 'r') as f:
    text = f.read()

text = text.replace(
    'type Lexer struct {\n\tinput        string\n\tposition     int\n\treadPosition int\n\tch           byte\n}',
    'type Lexer struct {\n\tinput        string\n\tposition     int\n\treadPosition int\n\tch           byte\n\tline         int\n\tcolumn       int\n}'
)
text = text.replace(
    'l := &Lexer{input: input}',
    'l := &Lexer{input: input, line: 1, column: 0}'
)

readChar_old = """func (l *Lexer) readChar() {
if l.readPosition >= len(l.input) {
l.ch = 0
} else {
l.ch = l.input[l.readPosition]
}
l.position = l.readPosition
l.readPosition++
}"""

readChar_new = """func (l *Lexer) readChar() {
if l.readPosition >= len(l.input) {
l.ch = 0
} else {
l.ch = l.input[l.readPosition]
}
l.position = l.readPosition
l.readPosition++
if l.ch == '\\n' {
l.line++
l.column = 0
} else {
l.column++
}
}"""
text = text.replace(readChar_old, readChar_new)

nextToken_old = """func (l *Lexer) NextToken() token.Token {
var tok token.Token
l.skipWhitespace()"""

nextToken_new = """func (l *Lexer) NextToken() token.Token {
var tok token.Token
l.skipWhitespace()

tokLine := l.line
if tokLine == 0 { tokLine = 1 }
tokCol := l.column
if tokCol == 0 { tokCol = 1 }"""
text = text.replace(nextToken_old, nextToken_new)

# handle newToken
text = text.replace('tok = newToken(token', 'tok = l.newToken(token')
text = text.replace('func newToken(tokenType token.TokenType, ch byte) token.Token {', 'func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {')

# Find token.Token{Type: ...} inside NextToken
# Replace with setting Line and Column manually or just add line col fields to all instantiations
text = text.replace('tok = token.Token{Type:', 'tok = token.Token{Line: tokLine, Column: tokCol, Type:')

# In literal scans, l.line, l.column advanced, so we just set them.
text = text.replace(
'tok.Type = token.GetTokenType(lit)\n\t\t\ttok.Literal = lit',
'tok.Type = token.GetTokenType(lit)\n\t\t\ttok.Literal = lit\n\t\t\ttok.Line = tokLine\n\t\t\ttok.Column = tokCol'
)

text = text.replace(
'tok.Type = token.INT\n\t\t\ttok.Literal = lit',
'tok.Type = token.INT\n\t\t\ttok.Literal = lit\n\t\t\ttok.Line = tokLine\n\t\t\ttok.Column = tokCol'
)
text = text.replace(
'tok.Type = token.STRING\n\t\ttok.Literal = l.readString()',
'tok.Type = token.STRING\n\t\ttok.Literal = l.readString()\n\t\ttok.Line = tokLine\n\t\ttok.Column = tokCol'
)
text = text.replace(
'tok.Type = token.EOF\n\t\ttok.Literal = ""',
'tok.Type = token.EOF\n\t\ttok.Literal = ""\n\t\ttok.Line = tokLine\n\t\ttok.Column = tokCol'
)

text = text.replace(
'return token.Token{Type: tokenType, Literal: string(ch)}',
'return token.Token{Type: tokenType, Literal: string(ch), Line: l.line, Column: l.column}'
)

with open('/media/nikhil/Windows/Nikium/lexer/lexer.go', 'w') as f:
    f.write(text)

