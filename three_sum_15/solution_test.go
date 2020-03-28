package three_sum_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "three sum",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "three sum",
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
			want: [][]int{
				{-2, 0, 2},
			},
		},
		{
			name: "three sum",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "three sum",
			args: args{
				nums: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
			},
			want: [][]int{
				{-5, 1, 4},
				{-4, 0, 4},
				{-4, 1, 3},
				{-2, -2, 4},
				{-2, 1, 1},
				{0, 0, 0},
			},
		},
		// -2, 0, 0, 0, 2, 2
		{
			name: "three sum",
			args: args{
				nums: []int{-2, 0, 0, 0, 2, 2},
			},
			want: [][]int{
				{-2, 0, 2},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, threeSum(tt.args.nums))
		})
	}
}
