package gen

import "fmt"

type StringEnum struct {
	Name   string
	Values []StringEnumValue
}

func (_ StringEnum) ConstantBaseType() string {
	return "string"
}

func (s StringEnum) ValuesTypeName() string {
	return fmt.Sprintf("%sValues", s.Name)
}

type StringEnumValue struct {
	Name         string
	Value        string
	PrintedValue string
	Comment      string
}

type IntEnum struct {
	Name   string
	Values []IntEnumValue
}

func (_ *IntEnum) ConstantBaseType() string {
	return "int"
}

type IntEnumValue struct {
	Name         string
	Value        int
	PrintedValue string
	Comment      string
}
