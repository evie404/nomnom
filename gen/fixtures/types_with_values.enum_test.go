package fixtures

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestIsCity(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"\"london\" is CityLondon",
			args{
				"london",
			},
			true,
		},
		{
			"\"oakland\" is CityOakland",
			args{
				"oakland",
			},
			true,
		},
		{
			"\"portland\" is CityPortland",
			args{
				"portland",
			},
			true,
		},
		{
			"\"seattle\" is CitySeattle",
			args{
				"seattle",
			},
			true,
		},
		{
			"\"San Francisco\" is CitySanFrancisco",
			args{
				"San Francisco",
			},
			true,
		},
		{
			"`\"` is CityQuotes",
			args{
				`"`,
			},
			true,
		},
		{
			"\"sekret\" is citySekret",
			args{
				"sekret",
			},
			true,
		},
		{
			"errors for a random number is not a City",
			args{
				strconv.Itoa(rand.Int()),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCity(tt.args.in); got != tt.want {
				t.Errorf("IsCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCity(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name  string
		args  args
		want  City
		want1 bool
	}{
		{
			"\"london\" is CityLondon",
			args{
				"london",
			},
			CityLondon,
			true,
		},
		{
			"\"oakland\" is CityOakland",
			args{
				"oakland",
			},
			CityOakland,
			true,
		},
		{
			"\"portland\" is CityPortland",
			args{
				"portland",
			},
			CityPortland,
			true,
		},
		{
			"\"seattle\" is CitySeattle",
			args{
				"seattle",
			},
			CitySeattle,
			true,
		},
		{
			"\"San Francisco\" is CitySanFrancisco",
			args{
				"San Francisco",
			},
			CitySanFrancisco,
			true,
		},
		{
			"`\"` is CityQuotes",
			args{
				`"`,
			},
			CityQuotes,
			true,
		},
		{
			"\"sekret\" is citySekret",
			args{
				"sekret",
			},
			citySekret,
			true,
		},
		{
			"errors for a random number is not a City",
			args{
				strconv.Itoa(rand.Int()),
			},
			City(""),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToCity(tt.args.in)
			if got != tt.want {
				t.Errorf("ToCity() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToCity() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToCityErr(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    City
		wantErr bool
	}{
		{
			"\"london\" is CityLondon",
			args{
				"london",
			},
			CityLondon,
			false,
		},
		{
			"\"oakland\" is CityOakland",
			args{
				"oakland",
			},
			CityOakland,
			false,
		},
		{
			"\"portland\" is CityPortland",
			args{
				"portland",
			},
			CityPortland,
			false,
		},
		{
			"\"seattle\" is CitySeattle",
			args{
				"seattle",
			},
			CitySeattle,
			false,
		},
		{
			"\"San Francisco\" is CitySanFrancisco",
			args{
				"San Francisco",
			},
			CitySanFrancisco,
			false,
		},
		{
			"`\"` is CityQuotes",
			args{
				`"`,
			},
			CityQuotes,
			false,
		},
		{
			"\"sekret\" is citySekret",
			args{
				"sekret",
			},
			citySekret,
			false,
		},
		{
			"errors for a random number is not a City",
			args{
				strconv.Itoa(rand.Int()),
			},
			City(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCityErr(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCityErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCityErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustToCity(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name      string
		args      args
		want      City
		wantPanic bool
	}{
		{
			"\"london\" is CityLondon",
			args{
				"london",
			},
			CityLondon,
			false,
		},
		{
			"\"oakland\" is CityOakland",
			args{
				"oakland",
			},
			CityOakland,
			false,
		},
		{
			"\"portland\" is CityPortland",
			args{
				"portland",
			},
			CityPortland,
			false,
		},
		{
			"\"seattle\" is CitySeattle",
			args{
				"seattle",
			},
			CitySeattle,
			false,
		},
		{
			"\"San Francisco\" is CitySanFrancisco",
			args{
				"San Francisco",
			},
			CitySanFrancisco,
			false,
		},
		{
			"`\"` is CityQuotes",
			args{
				`"`,
			},
			CityQuotes,
			false,
		},
		{
			"\"sekret\" is citySekret",
			args{
				"sekret",
			},
			citySekret,
			false,
		},
		{
			"panics for a random number is not a City",
			args{
				strconv.Itoa(rand.Int()),
			},
			City(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() { recover() }()
			}

			if got := MustToCity(tt.args.in); got != tt.want {
				t.Errorf("MustToCity() = %v, want %v", got, tt.want)
			}

			if tt.wantPanic {
				t.Errorf("did not panic")
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			true,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			true,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			true,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.in); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name  string
		args  args
		want  Number
		want1 bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			true,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			true,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			true,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToNumber(tt.args.in)
			if got != tt.want {
				t.Errorf("ToNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToNumberErr(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			false,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			false,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			false,
		},
		{
			"errors for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToNumberErr(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNumberErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToNumberErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustToNumber(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name      string
		args      args
		want      Number
		wantPanic bool
	}{
		{
			"1 is NumberOne",
			args{
				1,
			},
			NumberOne,
			false,
		},
		{
			"2 is NumberTwo",
			args{
				2,
			},
			NumberTwo,
			false,
		},
		{
			"3 is NumberThree",
			args{
				3,
			},
			NumberThree,
			false,
		},
		{
			"panics for a random number is not a Number",
			args{
				rand.Int(),
			},
			Number(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() { recover() }()
			}

			if got := MustToNumber(tt.args.in); got != tt.want {
				t.Errorf("MustToNumber() = %v, want %v", got, tt.want)
			}

			if tt.wantPanic {
				t.Errorf("did not panic")
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{
			"1 is NumberOne",
			args{
				"1",
			},
			NumberOne,
			false,
		},
		{
			"2 is NumberTwo",
			args{
				"2",
			},
			NumberTwo,
			false,
		},
		{
			"3 is NumberThree",
			args{
				"3",
			},
			NumberThree,
			false,
		},
		{
			"errors for non-numeric string",
			args{
				"abc",
			},
			Number(0),
			true,
		},
		{
			"errors for numeric string that is not a Number",
			args{
				strconv.Itoa(rand.Int()),
			},
			Number(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNumber(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
