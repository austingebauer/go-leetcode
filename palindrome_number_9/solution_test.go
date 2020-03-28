package palindrome_number_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "palindrome number",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "palindrome number",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "palindrome number",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "palindrome number",
			args: args{
				x: 1001,
			},
			want: true,
		},
		{
			name: "palindrome number",
			args: args{
				x: 1101,
			},
			want: false,
		},
		{
			name: "palindrome number",
			args: args{
				x: 1000030001,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.x))
		})
	}
}
