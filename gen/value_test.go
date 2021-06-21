package gen

import (
	"go/constant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalTypeAndValue(t *testing.T) {
	type args struct {
		expression string
	}
	type want struct {
		valueKind     constant.Kind
		stringVal     string
		extractString string
	}
	tests := []struct {
		name      string
		args      args
		want      want
		assertion assert.ErrorAssertionFunc
	}{
		{
			"string with backtick",
			args{
				"`test`",
			},
			want{
				constant.String,
				`"test"`,
				`"test"`,
			},
			assert.NoError,
		},
		{
			"string with backtick and double quote inside",
			args{
				"`test\"`",
			},
			want{
				constant.String,
				`"test\""`,
				`"test\""`,
			},
			assert.NoError,
		},
		{
			"string with double quotes",
			args{
				`"test"`,
			},
			want{
				constant.String,
				`"test"`,
				`"test"`,
			},
			assert.NoError,
		},
		{
			"integer",
			args{
				`123`,
			},
			want{
				constant.Int,
				`123`,
				`123`,
			},
			assert.NoError,
		},
		{
			"integer arithmetic",
			args{
				`1+1`,
			},
			want{
				constant.Int,
				`2`,
				`2`,
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evalTypeAndValue(tt.args.expression)
			tt.assertion(t, err)
			assert.Equal(t, tt.want.valueKind, got.Value.Kind())
			assert.Equal(t, tt.want.stringVal, got.Value.String())
			assert.Equal(t, tt.want.extractString, got.Value.ExactString())
		})
	}
}
