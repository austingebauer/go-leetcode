package two_sum_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "twoSum",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "twoSum",
			args: args{
				nums:   []int{2, 2, 11, 15},
				target: 4,
			},
			want: []int{0, 1},
		},
		{
			name: "twoSum",
			args: args{
				nums:   []int{2},
				target: 0,
			},
			want: []int{},
		},
		{
			name: "twoSum",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{},
		},
		{
			name: "twoSum",
			args: args{
				nums:   []int{2, 3, 11, 15},
				target: 17,
			},
			want: []int{0, 3},
		},
		{
			name: "twoSum",
			args: args{
				nums:   []int{11, 15, -1},
				target: 10,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}
}
