package fixtures

import (
	"math/rand"
	"strconv"
	"testing"
)

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
