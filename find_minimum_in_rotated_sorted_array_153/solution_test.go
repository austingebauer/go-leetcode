package find_minimum_in_rotated_sorted_array_153

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{
					3, 4, 5, 1, 2,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 2, 3,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMin(tt.args.nums))
			assert.Equal(t, tt.want, findMinLinear(tt.args.nums))
		})
	}
}
