package word_break_139

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "word break problem",
			args: args{
				s: "leetcode",
				wordDict: []string{
					"leet", "code",
				},
			},
			want: true,
		},
		{
			name: "word break problem",
			args: args{
				s: "applepenapple",
				wordDict: []string{
					"apple", "pen",
				},
			},
			want: true,
		},
		{
			name: "word break problem",
			args: args{
				s: "catsandog",
				wordDict: []string{
					"cats", "dog", "sand", "and", "cat",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, wordBreak(tt.args.s, tt.args.wordDict))
			assert.Equal(t, tt.want, wordBreak0(tt.args.s, tt.args.wordDict))
		})
	}
}
