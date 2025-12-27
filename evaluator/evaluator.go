package evaluator

import (
	"Nikium/ast"
	"fmt"
)

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

// --- Eval ---

func Eval(node ast.Node, env *Environment) Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return NULL

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &ReturnValue{Value: val}

	case *ast.PrintStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Println(val.Inspect())
		return NULL

	case *ast.IntegerLiteral:
		return &Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &String{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.BinaryExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfStatement:
		return evalIfExpression(node, env)

	case *ast.WhileStatement:
		return evalWhileStatement(node, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		return &Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        env,
		}

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	}

	return nil
}

// --- Indexing ---

func evalIndexExpression(left, index Object) Object {
	switch left := left.(type) {
	case *String:
		idx, ok := index.(*Integer)
		if !ok {
			return newError("index must be integer")
		}
		if idx.Value < 0 || idx.Value >= int64(len(left.Value)) {
			return NULL
		}
		return &String{Value: string(left.Value[idx.Value])}
	default:
		return newError("index operator not supported on %s", left.Type())
	}
}

// --- Helpers ---

func evalProgram(program *ast.Program, env *Environment) Object {
	var result Object
	for _, stmt := range program.Statements {
		result = Eval(stmt, env)
		switch r := result.(type) {
		case *ReturnValue:
			return r.Value
		case *Error:
			return r
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) Object {
	var result Object
	for _, stmt := range block.Statements {
		result = Eval(stmt, env)
		if result != nil {
			switch result.Type() {
			case RETURN_VALUE_OBJ, ERROR_OBJ:
				return result
			}
		}
	}
	return result
}

func evalWhileStatement(ws *ast.WhileStatement, env *Environment) Object {
	for {
		cond := Eval(ws.Condition, env)
		if isError(cond) {
			return cond
		}
		if !isTruthy(cond) {
			break
		}
		result := Eval(ws.Body, env)
		if result != nil {
			switch result.Type() {
			case RETURN_VALUE_OBJ, ERROR_OBJ:
				return result
			}
		}
	}
	return NULL
}

func evalIfExpression(ie *ast.IfStatement, env *Environment) Object {
	cond := Eval(ie.Condition, env)
	if isError(cond) {
		return cond
	}
	if isTruthy(cond) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}
	return NULL
}

func evalIdentifier(node *ast.Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError("identifier not found: %s", node.Value)
}

func evalExpressions(exps []ast.Expression, env *Environment) []Object {
	var result []Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		if fn.Native != nil {
			return fn.Native(args)
		}
		env := NewEnclosedEnvironment(fn.Env)
		for i, param := range fn.Parameters {
			env.Set(param.Value, args[i])
		}
		return unwrapReturnValue(Eval(fn.Body, env))
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func unwrapReturnValue(obj Object) Object {
	if rv, ok := obj.(*ReturnValue); ok {
		return rv.Value
	}
	return obj
}

// --- Boolean / Null ---

func nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isTruthy(obj Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// --- Prefix Expressions ---

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	if right.Type() != INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}
	return &Integer{Value: -right.(*Integer).Value}
}

// --- Infix Expressions ---

func evalInfixExpression(op string, left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return evalIntegerInfixExpression(op, left, right)

	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		lstr := left.(*String).Value
		rstr := right.(*String).Value

		if len(lstr) == 1 && len(rstr) == 1 {
			lv := int64(lstr[0])
			rv := int64(rstr[0])
			switch op {
			case "+":
				return &Integer{Value: lv + rv}
			case "-":
				return &Integer{Value: lv - rv}
			case "*":
				return &Integer{Value: lv * rv}
			case "/":
				return &Integer{Value: lv / rv}
			case "==":
				return nativeBoolToBooleanObject(lv == rv)
			case "!=":
				return nativeBoolToBooleanObject(lv != rv)
			case "<":
				return nativeBoolToBooleanObject(lv < rv)
			case ">":
				return nativeBoolToBooleanObject(lv > rv)
			default:
				return newError("unknown operator for chars: %s", op)
			}
		}

		if op == "+" {
			return &String{Value: lstr + rstr}
		}
		return newError("unknown operator for strings: %s", op)

	case left.Type() == STRING_OBJ && right.Type() == INTEGER_OBJ:
		lstr := left.(*String).Value
		if len(lstr) != 1 {
			return newError("cannot perform arithmetic on multi-char string")
		}
		lv := int64(lstr[0])
		rv := right.(*Integer).Value
		switch op {
		case "+":
			return &Integer{Value: lv + rv}
		case "-":
			return &Integer{Value: lv - rv}
		case "*":
			return &Integer{Value: lv * rv}
		case "/":
			return &Integer{Value: lv / rv}
		default:
			return newError("unknown operator for char and int: %s", op)
		}

	case left.Type() == INTEGER_OBJ && right.Type() == STRING_OBJ:
		rstr := right.(*String).Value
		if len(rstr) != 1 {
			return newError("cannot perform arithmetic on multi-char string")
		}
		lv := left.(*Integer).Value
		rv := int64(rstr[0])
		switch op {
		case "+":
			return &Integer{Value: lv + rv}
		case "-":
			return &Integer{Value: lv - rv}
		case "*":
			return &Integer{Value: lv * rv}
		case "/":
			return &Integer{Value: lv / rv}
		default:
			return newError("unknown operator for int and char: %s", op)
		}

	case op == "==":
		return nativeBoolToBooleanObject(left == right)
	case op == "!=":
		return nativeBoolToBooleanObject(left != right)

	default:
		return newError("type mismatch: %s %s %s", left.Type(), op, right.Type())
	}
}

func evalIntegerInfixExpression(op string, left, right Object) Object {
	lv := left.(*Integer).Value
	rv := right.(*Integer).Value
	switch op {
	case "+":
		return &Integer{Value: lv + rv}
	case "-":
		return &Integer{Value: lv - rv}
	case "*":
		return &Integer{Value: lv * rv}
	case "/":
		return &Integer{Value: lv / rv}
	case "<":
		return nativeBoolToBooleanObject(lv < rv)
	case ">":
		return nativeBoolToBooleanObject(lv > rv)
	case "==":
		return nativeBoolToBooleanObject(lv == rv)
	case "!=":
		return nativeBoolToBooleanObject(lv != rv)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
	}
}

// --- Utility ---

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	return obj != nil && obj.Type() == ERROR_OBJ
}
