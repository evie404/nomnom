package gen

import (
	"fmt"
	"strings"
)

type Options struct {
	GenerateValuesField  bool
	GenerateValuesStruct bool
}

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

func (e Enum) RandValue() string {
	switch e.BaseType {
	case "string":
		return "strconv.Itoa(rand.Int())"
	case "int":
		return "rand.Int()"
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

func (e EnumValue) RawValue() string {
	rawValue, _ := EvalValueString(e.Value)
	return trimQuotes(rawValue)
}

func trimQuotes(s string) string {
	if len(s) < 2 {
		return s
	}

	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return s[1 : len(s)-1]
	}

	return s
}

func (e EnumValue) EscapedValue() string {
	return strings.Replace(e.Value, `"`, `\"`, -1)
}
