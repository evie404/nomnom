package gen

import (
	"fmt"
)

func GenerateEnumHelpers(pkgName string, enums []Enum) ([]byte, error) {
	var result []byte

	for _, enum := range enums {
		conversions, err := ConversionsTemplate(enum)
		if err != nil {
			return nil, fmt.Errorf("generating conversions: %w", err)
		}

		valuesStruct, err := ValuesStructTemplate(enum)
		if err != nil {
			return nil, fmt.Errorf("generating values struct: %w", err)
		}

		result = append(result, conversions...)
		result = append(result, "\n"[0])
		result = append(result, valuesStruct...)
	}

	return formatCode(pkgName, result)
}
