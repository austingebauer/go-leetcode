package add_two_numbers_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "same length",
			args: args{
				word1: "abc",
				word2: "pqr",
			},
			want: "apbqcr",
		},
		{
			name: "different length 1",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "different length 2",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeAlternately(tt.args.word1, tt.args.word2), "mergeAlternately(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}
