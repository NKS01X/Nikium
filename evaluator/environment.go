package evaluator

import (
	"fmt"
)

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	env := &Environment{store: s, outer: nil}
	env.Set("readchar", &Function{
		Native: nativeReadChar,
	})
	env.Set("len", &Function{
		Native: func(args []Object) Object {
			if len(args) != 1 {
				return &Error{
					Message: fmt.Sprintf("len: expected 1 argument, got %d", len(args)),
				}
			}
			switch arg := args[0].(type) {
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return &Error{
					Message: fmt.Sprintf("len: unsupported type %s", arg.Type()),
				}
			}
		},
	})
	env.Set("readline", &Function{
		Native: nativeReadLine,
	})

	env.Set("Print", &Function{
		Native: func(args []Object) Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect(), " ")
			}
			fmt.Println()
			return NULL
		},
	})

	env.Set("push", &Function{
		Native: func(args []Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("push: expected 2 arguments, got %d", len(args))}
			}
			arr, ok := args[0].(*Array)
			if !ok {
				return &Error{Message: "push: first argument must be an array"}
			}
			newElements := make([]Object, len(arr.Elements)+1)
			copy(newElements, arr.Elements)
			newElements[len(arr.Elements)] = args[1]
			return &Array{Elements: newElements}
		},
	})

	return env
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
