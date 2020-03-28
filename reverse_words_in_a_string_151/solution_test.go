package reverse_words_in_a_string_151

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse words in a string",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "reverse words in a string",
			args: args{
				s: "  hello world!  ",
			},
			want: "world! hello",
		},
		{
			name: "reverse words in a string",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseWords(tt.args.s))
			assert.Equal(t, tt.want, reverseWordsArrayReversal(tt.args.s))
		})
	}
}
