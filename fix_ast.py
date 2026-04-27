import re

with open('/media/nikhil/Windows/Nikium/token/token.go', 'r') as f:
    text = f.read()

text = text.replace(
    'type Token struct {\n\tType    TokenType\n\tLiteral string\n}',
    'type Token struct {\n\tType    TokenType\n\tLiteral string\n\tLine    int\n\tColumn  int\n}'
)

with open('/media/nikhil/Windows/Nikium/token/token.go', 'w') as f:
    f.write(text)

with open('/media/nikhil/Windows/Nikium/ast/ast.go', 'r') as f:
    text = f.read()

text = text.replace(
    'type Node interface {\n\tTokenLiteral() string\n\tString() string\n}',
    'type Node interface {\n\tTokenLiteral() string\n\tString() string\n\tGetToken() token.Token\n}'
)

# find all types and add GetToken func
types = re.findall(r'type (\w+) struct \{', text)

lines = []
for t in types:
    if t == 'Program':
        lines.append(f'func (n *{t}) GetToken() token.Token {{ if len(n.Statements) > 0 {{ return n.Statements[0].GetToken() }}; return token.Token{{}} }}')
    elif t == 'IndexExpression':
        lines.append(f'func (n *{t}) GetToken() token.Token {{ return n.Left.GetToken() }}')
    else:
        lines.append(f'func (n *{t}) GetToken() token.Token {{ return n.Token }}')

text += '\n\n' + '\n'.join(lines) + '\n'

with open('/media/nikhil/Windows/Nikium/ast/ast.go', 'w') as f:
    f.write(text)

