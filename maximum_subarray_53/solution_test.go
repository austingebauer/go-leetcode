package maximum_subarray_53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contiguous subarray which has the largest sum",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "contiguous subarray which has the largest sum",
			args: args{
				nums: []int{-1000, 1000, 1},
			},
			want: 1001,
		},
		{
			name: "contiguous subarray which has the largest sum",
			args: args{
				nums: []int{-1, 2, -1, -3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSubArray(tt.args.nums))
			assert.Equal(t, tt.want, maxSubArray2(tt.args.nums))
		})
	}
}
