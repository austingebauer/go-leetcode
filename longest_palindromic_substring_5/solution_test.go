package longest_palindromic_substring_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "largest palindrome in string",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "b",
			},
			want: "b",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "baaab",
			},
			want: "baaab",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPalindrome(tt.args.s))
		})
	}
}

func Test_longestPalindrome0(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "largest palindrome in string",
			args: args{
				s: "babad",
			},
			want: "aba",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "b",
			},
			want: "b",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "baaab",
			},
			want: "baaab",
		},
		{
			name: "largest palindrome in string",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPalindrome0(tt.args.s))
		})
	}
}
