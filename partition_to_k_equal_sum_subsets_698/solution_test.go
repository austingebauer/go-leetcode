package partition_to_k_equal_sum_subsets_698

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					1, 1, 2, 2,
				},
				k: 3,
			},
			want: true,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					4, 3, 2,
				},
				k: 4,
			},
			want: false,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					4, 3, 2, 3,
				},
				k: 4,
			},
			want: false,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					1, 1, 1, 1,
				},
				k: 4,
			},
			want: true,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					1, 1, 1, 2,
				},
				k: 3,
			},
			want: false,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					4, 3, 2, 3, 5, 2, 1,
				},
				k: 2,
			},
			want: true,
		},
		{
			name: "can partition into k subsets",
			args: args{
				nums: []int{
					10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6,
				},
				k: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canPartitionKSubsets(tt.args.nums, tt.args.k))
		})
	}
}
