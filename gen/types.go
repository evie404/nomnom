package gen

import (
	"fmt"
	"strings"
)

type Enum struct {
	Name     string
	BaseType string
	Values   []EnumValue
}

func (e Enum) ConstantBaseType() string {
	return e.BaseType
}

func (e Enum) ValuesTypeName() string {
	return fmt.Sprintf("%sValues", e.Name)
}

func (e Enum) TypeName() string {
	return e.Name
}

func (e Enum) InputVarName() string {
	return "in"
}

func (e Enum) NullValue() string {
	switch e.BaseType {
	case "string":
		return `""`
	case "int":
		return "0"
	}

	return ""
}

func (e Enum) VarName() string {
	return strings.ToLower(e.TypeName())
}

type EnumValue struct {
	Name    string
	Value   string
	Comment string
}
