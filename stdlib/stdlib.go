package stdlib

import (
	"Nikium/evaluator"
	"fmt"
)

func Register(env *evaluator.Environment) {
	env.Set("abs", &evaluator.Function{
		Native: func(args []evaluator.Object) evaluator.Object {
			if len(args) != 1 {
				return &evaluator.Error{
					Message: fmt.Sprintf("abs: expected 1 argument, got %d", len(args)),
				}
			}
			arg, ok := args[0].(*evaluator.Integer)
			if !ok {
				return &evaluator.Error{
					Message: fmt.Sprintf("abs: expected an integer, got %s", args[0].Type()),
				}
			}
			if arg.Value < 0 {
				return &evaluator.Integer{Value: -arg.Value}
			}
			return arg
		},
	})
}
