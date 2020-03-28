package longest_common_subsequence_1143

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest common subsequence",
			args: args{
				text1: "bsbininm",
				text2: "jmjkbkjkv",
			},
			want: 1,
		},
		{
			name: "longest common subsequence",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "longest common subsequence",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "longest common subsequence",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonSubsequence(tt.args.text1, tt.args.text2))
			assert.Equal(t, tt.want, longestCommonSubsequence2(tt.args.text1, tt.args.text2))
		})
	}
}
