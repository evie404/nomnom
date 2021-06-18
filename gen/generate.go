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

		result = append(result, conversions...)
		result = append(result, "\n"[0])

		if enum.BaseType == "int" {
			numConversions, err := NumericConversionsTemplate(enum)
			if err != nil {
				return nil, fmt.Errorf("generating numeric conversions: %w", err)
			}

			result = append(result, numConversions...)
			result = append(result, "\n"[0])
		}

		valuesStruct, err := ValuesStructTemplate(enum)
		if err != nil {
			return nil, fmt.Errorf("generating values struct: %w", err)
		}

		result = append(result, valuesStruct...)
	}

	return formatCode(pkgName, nil, result)
}

func GenerateEnumHelpersTests(pkgName string, enums []Enum) ([]byte, error) {
	var result []byte

	for _, enum := range enums {
		conversions, err := ConversionsTestTemplate(enum)
		if err != nil {
			return nil, fmt.Errorf("generating conversions: %w", err)
		}

		result = append(result, conversions...)
		result = append(result, "\n"[0])

		if enum.BaseType == "int" {
			numConversions, err := NumericConversionsTestTemplate(enum)
			if err != nil {
				return nil, fmt.Errorf("generating numeric conversions: %w", err)
			}

			result = append(result, numConversions...)
			result = append(result, "\n"[0])
		}
	}

	return formatCode(pkgName, []string{"math/rand"}, result)
}
