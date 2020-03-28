package top_k_frequent_elements_347

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "top k frequent elements",
			args: args{
				nums: []int{
					1, 1, 1, 2, 2, 3,
				},
				k: 2,
			},
			want: []int{
				1, 2,
			},
		},
		{
			name: "top k frequent elements",
			args: args{
				nums: []int{
					1,
				},
				k: 1,
			},
			want: []int{
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, topKFrequent(tt.args.nums, tt.args.k))
		})
	}
}
