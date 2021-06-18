package gen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateEnumHelpers(t *testing.T) {
	type args struct {
		pkgName string
		enums   []Enum
		opts    Options
	}
	tests := []struct {
		name            string
		args            args
		wantFixturePath string
		assertion       assert.ErrorAssertionFunc
	}{
		{
			"",
			args{
				pkgName: "fixtures",
				enums: []Enum{
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
						Values: []EnumValue{
							{
								Name:  "NumberOne",
								Value: "1",
							},
							{
								Name:  "NumberTwo",
								Value: "2",
							},
							{
								Name:  "NumberThree",
								Value: "3",
							},
						},
					},
				},
				opts: Options{
					GenerateValuesStruct: true,
					GenerateValuesField:  true,
				},
			},
			filepath.Join("fixtures", "types_with_values.enum.go"),
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.wantFixturePath)
			require.NoError(t, err)

			got, err := GenerateEnumHelpers(tt.args.pkgName, tt.args.enums, tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, string(want), string(got))
		})
	}
}

func TestGenerateEnumHelpersTests(t *testing.T) {
	type args struct {
		pkgName string
		enums   []Enum
	}
	tests := []struct {
		name            string
		args            args
		wantFixturePath string
		assertion       assert.ErrorAssertionFunc
	}{
		{
			"",
			args{
				pkgName: "fixtures",
				enums: []Enum{
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
						Values: []EnumValue{
							{
								Name:  "NumberOne",
								Value: "1",
							},
							{
								Name:  "NumberTwo",
								Value: "2",
							},
							{
								Name:  "NumberThree",
								Value: "3",
							},
						},
					},
				},
			},
			filepath.Join("fixtures", "types_with_values.enum_test.go"),
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.wantFixturePath)
			require.NoError(t, err)

			got, err := GenerateEnumHelpersTests(tt.args.pkgName, tt.args.enums)
			tt.assertion(t, err)
			assert.Equal(t, string(want), string(got))
		})
	}
}
