package gen

import (
	"fmt"

	"golang.org/x/tools/imports"
)

func GenerateEnumHelpers(pkgName string, strEnum StringEnum) ([]byte, error) {
	conversions, err := ConversionsTemplate(strEnum)
	if err != nil {
		return nil, fmt.Errorf("generating conversions: %w", err)
	}

	valuesStruct, err := ValuesStructTemplate(strEnum)
	if err != nil {
		return nil, fmt.Errorf("generating values struct: %w", err)
	}

	result := make([]byte, 0, len(conversions)+len(valuesStruct))
	result = append(result, []byte("package "+pkgName)...)
	result = append(result, "\n"[0])
	result = append(result, conversions...)
	result = append(result, "\n"[0])
	result = append(result, valuesStruct...)

	result, err = imports.Process("", result, nil)
	if err != nil {
		return nil, fmt.Errorf("running goimports: %w", err)
	}

	return result, nil
}
