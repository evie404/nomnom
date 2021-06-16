package gen

type StringEnum struct {
	Name   string
	Values []StringEnumValue
}

type StringEnumValue struct {
	Name    string
	Value   string
	Comment string
}

type IntEnum struct {
	Name   string
	Values []IntEnumValue
}

type IntEnumValue struct {
	Name    string
	Value   int
	Comment string
}
