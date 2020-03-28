package house_robber_198

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "house robber single number case",
			args: args{
				nums: []int{
					1,
				},
			},
			want: 1,
		},
		{
			name: "house robber two numbers case",
			args: args{
				nums: []int{
					1, 2,
				},
			},
			want: 2,
		},
		{
			name: "house robber three numbers case [0 + 2]",
			args: args{
				nums: []int{
					1, 2, 3,
				},
			},
			want: 4,
		},
		{
			name: "house robber three numbers case [1]",
			args: args{
				nums: []int{
					1, 3, 1,
				},
			},
			want: 3,
		},
		{
			name: "house robber",
			args: args{
				nums: []int{
					2, 7, 9, 3, 1,
				},
			},
			want: 12,
		},
		{
			name: "house robber",
			args: args{
				nums: []int{
					6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3,
				},
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rob(tt.args.nums))
			assert.Equal(t, tt.want, rob0(tt.args.nums))
		})
	}
}
