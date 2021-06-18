package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"golang.org/x/tools/imports"
)

func ValuesStructTemplate(enum Enum) ([]byte, error) {
	return runTemplate(filepath.Join("templates", "values_struct.go.tmpl"), enum)
}

func ConversionsTemplate(enum Enum) ([]byte, error) {
	return runTemplate(filepath.Join("templates", "conversions.go.tmpl"), enum)
}

func NumericConversionsTemplate(enum Enum) ([]byte, error) {
	return runTemplate(filepath.Join("templates", "numeric_conversions.go.tmpl"), enum)
}

func runTemplate(templatePath string, enum Enum) ([]byte, error) {
	rawTemplate, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("reading template file: %w", err)
	}

	t, err := template.New("letter").Parse(string(rawTemplate))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	var b bytes.Buffer

	err = t.Execute(&b, enum)
	if err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return b.Bytes(), nil
}

func formatCode(pkgName string, content []byte) ([]byte, error) {
	result := []byte("package " + pkgName)
	result = append(result, "\n"[0])
	result = append(result, content...)

	formatted, err := imports.Process("", result, nil)
	if err != nil {
		return nil, fmt.Errorf("running goimports: %w", err)
	}

	return formatted, nil
}
