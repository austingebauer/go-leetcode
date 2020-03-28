package palindromic_substrings_647

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "palindromic substrings",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "palindromic substrings",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
		{
			name: "palindromic substrings",
			args: args{
				s: "aaaaa",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countSubstrings(tt.args.s))
		})
	}
}
