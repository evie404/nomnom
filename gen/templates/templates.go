package templates

import (
	_ "embed"
)

var (
	//go:embed conversions_test.go.tmpl
	ConversionsTest []byte

	//go:embed conversions.go.tmpl
	Conversions []byte

	//go:embed numeric_conversions.go.tmpl
	NumericConversions []byte

	//go:embed numeric_conversions_test.go.tmpl
	NumericConversionsTest []byte

	//go:embed values_struct.go.tmpl
	ValuesStruct []byte

	//go:embed values_field.go.tmpl
	ValuesField []byte
)
