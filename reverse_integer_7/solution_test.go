package reverse_integer_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reverse integer",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "reverse integer",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "reverse integer",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "reverse integer",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "reverse integer with positive overflow",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
		{
			name: "reverse integer with negative overflow",
			args: args{
				x: -2147483648,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverse(tt.args.x))
		})
	}
}
