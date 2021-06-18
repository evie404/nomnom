package gen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValuesStructTemplate(t *testing.T) {
	type args struct {
		enum Enum
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
				enum: Enum{
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
			},
			filepath.Join("templates", "fixtures", "string_values_struct.go"),
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.wantFixturePath)
			require.NoError(t, err)

			got, err := ValuesStructTemplate(tt.args.enum)
			tt.assertion(t, err)

			formattedGot, err := formatCode("fixtures", got)
			require.NoError(t, err)

			assert.Equal(t, string(want), string(formattedGot))
		})
	}
}

func TestConversionsTemplate(t *testing.T) {
	type args struct {
		enum Enum
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
				enum: Enum{
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
			},
			filepath.Join("templates", "fixtures", "string_conversions.go"),
			assert.NoError,
		},
		{
			"",
			args{
				enum: Enum{
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
			filepath.Join("templates", "fixtures", "int_conversions.go"),
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.wantFixturePath)
			require.NoError(t, err)

			got, err := ConversionsTemplate(tt.args.enum)
			tt.assertion(t, err)

			formattedGot, err := formatCode("fixtures", got)
			require.NoError(t, err)

			assert.Equal(t, string(want), string(formattedGot))
		})
	}
}

func TestNumericConversionsTemplate(t *testing.T) {
	type args struct {
		enum Enum
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
				enum: Enum{
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
			filepath.Join("templates", "fixtures", "int_numeric_conversions.go"),
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := ioutil.ReadFile(tt.wantFixturePath)
			require.NoError(t, err)

			got, err := NumericConversionsTemplate(tt.args.enum)
			tt.assertion(t, err)

			formattedGot, err := formatCode("fixtures", got)
			require.NoError(t, err)

			assert.Equal(t, string(want), string(formattedGot))
		})
	}
}
