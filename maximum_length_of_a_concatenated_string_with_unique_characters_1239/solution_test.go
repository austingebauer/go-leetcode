package maximum_length_of_a_concatenated_string_with_unique_characters_1239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"un", "iq", "ue",
				},
			},
			want: 4,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"cha", "r", "act", "ers",
				},
			},
			want: 6,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"abcdefghijklmnopqrstuvwxyz",
				},
			},
			want: 26,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"co", "dil", "ity",
				},
			},
			want: 5,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"abc", "kkk", "def", "csv",
				},
			},
			want: 6,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"abc", "ade", "akl",
				},
			},
			want: 3,
		},
		{
			name: "maximum length of a concatenated string with unique characters",
			args: args{
				arr: []string{
					"ab", "ba", "cd", "dc", "ef", "fe", "gh", "hg", "ij", "ji", "kl", "lk", "mn", "nm", "op", "po",
				},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxLength(tt.args.arr))
		})
	}
}
