package valid_palindrome_125

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid palindrome",
			args: args{
				s: "racecar",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "ab",
			},
			want: false,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "a:a",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "::",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: " ",
			},
			want: true,
		},
		{
			name: "valid palindrome",
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.s))
		})
	}
}
