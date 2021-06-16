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
		want     []*StringEnum
		want1    []*IntEnum
	}{
		{
			"list enum types",
			filepath.Join("fixtures", "types.go"),
			[]*StringEnum{
				{
					Name: "StringEnumType1",
				},
				{
					Name: "StringEnumType2",
				},
				{
					Name: "StringEnumType3",
				},
				{
					Name: "StringEnumType4",
				},
				{
					Name: "StringEnumType5",
				},
			},
			[]*IntEnum{
				{
					Name: "IntEnumType1",
				},
				{
					Name: "IntEnumType2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got, got1 := listEnumTypes(astFile.Decls)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_listEnumValues(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     map[string][]StringEnumValue
		want1    map[string][]IntEnumValue
	}{
		{
			"list string enum types",
			filepath.Join("fixtures", "types_with_values.go"),
			map[string][]StringEnumValue{
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
				},
			},
			map[string][]IntEnumValue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got, got1 := listEnumValues(astFile.Decls)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestListEnumsTypesValues(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     []*StringEnum
		want1    []*IntEnum
	}{
		{
			"list string enum types",
			filepath.Join("fixtures", "types_with_values.go"),
			[]*StringEnum{
				{
					Name: "City",
					Values: []StringEnumValue{
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
					},
				},
			},
			[]*IntEnum{
				{
					Name: "Number",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astFile, err := readFileAst(tt.filepath)
			require.NoError(t, err)

			got, got1 := ListEnumsTypesValues(astFile.Decls)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
