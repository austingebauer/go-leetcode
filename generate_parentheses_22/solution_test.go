package generate_parentheses_22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "generate parentheses",
			args: args{
				n: 0,
			},
			want: []string{
				"",
			},
		},
		{
			name: "generate parentheses",
			args: args{
				n: 1,
			},
			want: []string{
				"()",
			},
		},
		{
			name: "generate parentheses",
			args: args{
				n: 2,
			},
			want: []string{
				"(())",
				"()()",
			},
		},
		{
			name: "generate parentheses",
			args: args{
				n: 3,
			},
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, generateParenthesis(tt.args.n))
		})
	}
}
