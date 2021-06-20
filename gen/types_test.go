package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trimQuotes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"123"},
			"123",
		},
		{
			args{"\"123"},
			"\"123",
		},
		{
			args{"\"abc\""},
			"abc",
		},
		{
			args{`abs`},
			"abs",
		},
		{
			args{`abs"`},
			"abs\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			assert.Equal(t, tt.want, trimQuotes(tt.args.s))
		})
	}
}
