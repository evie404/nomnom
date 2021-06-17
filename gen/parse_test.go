package gen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_listEnumTypes(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     []*Enum
	}{
		{
			"list enum types",
			filepath.Join("fixtures", "types.go"),
			[]*Enum{
				{
					Name:     "IntEnumType1",
					BaseType: "int",
				},
				{
					Name:     "IntEnumType2",
					BaseType: "int",
				},
				{
					Name:     "StringEnumType1",
					BaseType: "string",
				},
				{
					Name:     "StringEnumType2",
					BaseType: "string",
				},
				{
					Name:     "StringEnumType3",
					BaseType: "string",
				},
				{
					Name:     "StringEnumType4",
					BaseType: "string",
				},
				{
					Name:     "StringEnumType5",
					BaseType: "string",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got := listEnumTypes(astFile.Decls)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_listEnumValues(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     map[string][]EnumValue
	}{
		{
			"list string enum types",
			filepath.Join("fixtures", "types_with_values.go"),
			map[string][]EnumValue{
				"City": {
					{
						Name:  "CityLondon",
						Value: "\"london\"",
					},
					{
						Name:  "CityOakland",
						Value: "\"oakland\"",
					},
					{
						Name:  "CityPortland",
						Value: "\"portland\"",
					},
					{
						Name:  "CitySeattle",
						Value: "\"seattle\"",
					},
					{
						Name:  "CitySanFrancisco",
						Value: "\"San Francisco\"",
					},
					{
						Name:  "CityQuotes",
						Value: "`\"`",
					},
					{
						Name:  "citySekret",
						Value: "\"sekret\"",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got := listEnumValues(astFile.Decls)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListEnumsTypesValues(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     []*Enum
	}{
		{
			"list string enum types",
			filepath.Join("fixtures", "types_with_values.go"),
			[]*Enum{
				{
					Name:     "City",
					BaseType: "string",
					Values: []EnumValue{
						{
							Name:  "CityLondon",
							Value: "\"london\"",
						},
						{
							Name:  "CityOakland",
							Value: "\"oakland\"",
						},
						{
							Name:  "CityPortland",
							Value: "\"portland\"",
						},
						{
							Name:  "CitySeattle",
							Value: "\"seattle\"",
						},
						{
							Name:  "CitySanFrancisco",
							Value: "\"San Francisco\"",
						},
						{
							Name:  "CityQuotes",
							Value: "`\"`",
						},
						{
							Name:  "citySekret",
							Value: "\"sekret\"",
						},
					},
				},
				{
					Name:     "Number",
					BaseType: "int",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got := ListEnumsTypesValues(astFile.Decls)
			assert.Equal(t, tt.want, got)
		})
	}
}
