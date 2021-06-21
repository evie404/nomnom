package gen

import (
	"go/token"
	"go/types"
)

func EvalValueString(expression string) (string, error) {
	typeAndValue, err := evalTypeAndValue(expression)
	if err != nil {
		return "", err
	}

	return typeAndValue.Value.String(), nil
}

func evalTypeAndValue(expression string) (types.TypeAndValue, error) {
	fset := token.NewFileSet()
	return types.Eval(fset, nil, 0, expression)
}
