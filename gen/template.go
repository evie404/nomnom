package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

func ValuesStructTemplate(strEnum StringEnum) ([]byte, error) {
	rawTemplate, err := ioutil.ReadFile(filepath.Join("templates", "values_struct.go.tmpl"))
	if err != nil {
		return nil, fmt.Errorf("reading template file: %w", err)
	}

	t, err := template.New("letter").Parse(string(rawTemplate))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	var b bytes.Buffer

	err = t.Execute(&b, strEnum)
	if err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return b.Bytes(), nil
}
