package verifying_an_alien_dictionary_953

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "verify alien dictionary",
			args: args{
				words: []string{
					"hello", "leetcode",
				},
				order: "hlabcdefgijkmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "verify alien dictionary",
			args: args{
				words: []string{
					"word", "world", "row",
				},
				order: "worldabcefghijkmnpqstuvxyz",
			},
			want: false,
		},
		{
			name: "verify alien dictionary",
			args: args{
				words: []string{
					"apple", "app",
				},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAlienSorted(tt.args.words, tt.args.order))
		})
	}
}
