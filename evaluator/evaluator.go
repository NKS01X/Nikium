package evaluator

import (
	"Nikium/ast"
	"fmt"
)

// Eval walks the AST and evaluates nodes
func Eval(node ast.Node, env *Environment) Object {
	switch n := node.(type) {

	// ---------------- Statements ----------------
	case *ast.Program:
		var result Object
		for _, stmt := range n.Statements {
			result = Eval(stmt, env)
		}
		return result

	case *ast.LetStatement:
		val := Eval(n.Value, env)
		env.Set(n.Name.Value, val)
		return val

	case *ast.PrintStatement:
		val := Eval(n.Value, env)
		fmt.Println(val.Inspect())
		return val

	case *ast.BlockStatement:
		var result Object
		for _, stmt := range n.Statements {
			result = Eval(stmt, env)
		}
		return result

	case *ast.IfStatement:
		cond := Eval(n.Condition, env)
		if isTruthy(cond) {
			return Eval(n.Consequence, env)
		} else if n.Alternative != nil {
			return Eval(n.Alternative, env)
		}
		return &Null{}

	case *ast.WhileStatement:
		var result Object
		for {
			cond := Eval(n.Condition, env)
			if !isTruthy(cond) {
				break
			}
			result = Eval(n.Body, env)
		}
		return result

	// ---------------- Expressions ----------------
	case *ast.IntegerLiteral:
		if n.Type == "i32" {
			return &Integer{Value: n.Value, IntType: I32_OBJ}
		}
		return &Integer{Value: n.Value, IntType: I64_OBJ}

	case *ast.StringLiteral:
		return &String{Value: n.Value}

	case *ast.Identifier:
		if val, ok := env.Get(n.Value); ok {
			return val
		}
		return newError("identifier not found: " + n.Value)

	case *ast.BinaryExpression:
		left := Eval(n.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(n.Right, env)
		if isError(right) {
			return right
		}
		return evalBinaryExpression(n.Operator, left, right)
	}

	return nil
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == "ERROR"
	}
	return false
}

// ---------------- Helpers ----------------

func isTruthy(obj Object) bool {
	switch o := obj.(type) {
	case *Boolean:
		return o.Value
	case *Null:
		return false
	case *Integer:
		return o.Value != 0
	case *String:
		return o.Value != ""
	default:
		return false
	}
}

func evalBinaryExpression(op string, left, right Object) Object {
	switch {
	case (left.Type() == I32_OBJ || left.Type() == I64_OBJ) && (right.Type() == I32_OBJ || right.Type() == I64_OBJ):
		l := left.(*Integer).Value
		r := right.(*Integer).Value
		switch op {
		case "+":
			return &Integer{Value: l + r, IntType: I64_OBJ}
		case "-":
			return &Integer{Value: l - r, IntType: I64_OBJ}
		case "*":
			return &Integer{Value: l * r, IntType: I64_OBJ}
		case "==":
			return &Boolean{Value: l == r}
		case "!=":
			return &Boolean{Value: l != r}
		case "<":
			return &Boolean{Value: l < r}
		case ">":
			return &Boolean{Value: l > r}
		default:
			return &Error{Message: fmt.Sprintf("unknown operator: %s %s %s", left.Type(), op, right.Type())}
		}
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		if op == "+" {
			return &String{Value: left.(*String).Value + right.(*String).Value}
		}
		return &Error{Message: fmt.Sprintf("unknown operator: %s %s %s", left.Type(), op, right.Type())}
	}
	return &Error{Message: fmt.Sprintf("type mismatch: %s %s %s", left.Type(), op, right.Type())}
}
