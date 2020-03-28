package first_unique_character_in_a_string_387

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first unique character in string",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "first unique character in string",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, firstUniqChar(tt.args.s))
		})
	}
}
