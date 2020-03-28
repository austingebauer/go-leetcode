package maximum_product_subarray_152

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum product subarray",
			args: args{
				nums: []int{
					2, 3, -2, 4,
				},
			},
			want: 6,
		},
		{
			name: "maximum product subarray",
			args: args{
				nums: []int{
					-2,
				},
			},
			want: -2,
		},
		{
			name: "maximum product subarray",
			args: args{
				nums: []int{
					-2, 0, -1,
				},
			},
			want: 0,
		},
		{
			name: "maximum product subarray",
			args: args{
				nums: []int{
					-4, -3,
				},
			},
			want: 12,
		},
		{
			name: "maximum product subarray",
			args: args{
				nums: []int{
					-3, 0, 1, -2,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProduct(tt.args.nums))
		})
	}
}
