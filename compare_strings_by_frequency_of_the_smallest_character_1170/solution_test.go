package compare_strings_by_frequency_of_the_smallest_character_1170

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "compare strings by frequency of the smallest character",
			args: args{
				queries: []string{
					"cbd",
				},
				words: []string{
					"zaaaz",
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "compare strings by frequency of the smallest character",
			args: args{
				queries: []string{
					"bbb", "cc",
				},
				words: []string{
					"a", "aa", "aaa", "aaaa",
				},
			},
			want: []int{
				1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numSmallerByFrequency(tt.args.queries, tt.args.words))
		})
	}
}
