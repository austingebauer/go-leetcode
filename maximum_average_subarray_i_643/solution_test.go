package maximum_average_subarray_i_643

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "many numbers",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75000,
		},
		{
			name: "single number",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMaxAverage(tt.args.nums, tt.args.k))
		})
	}
}
