package longest_substring_without_repeating_characters_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest substring without repeating characters",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "longest substring without repeating characters",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "longest substring without repeating characters",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "longest substring without repeating characters",
			args: args{
				s: "adcab",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.args.s))
		})
	}
}
