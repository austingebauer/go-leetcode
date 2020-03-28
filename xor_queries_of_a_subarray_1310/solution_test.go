package xor_queries_of_a_subarray_1310

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	type args struct {
		arr     []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "xor queries",
			args: args{
				arr: []int{
					1, 3, 4, 8,
				},
				queries: [][]int{
					{0, 1}, {1, 2}, {0, 3}, {3, 3},
				},
			},
			want: []int{
				2, 7, 14, 8,
			},
		},
		{
			name: "xor queries",
			args: args{
				arr: []int{
					4, 8, 2, 10,
				},
				queries: [][]int{
					{2, 3}, {1, 3}, {0, 0}, {0, 3},
				},
			},
			want: []int{
				8, 0, 4, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, xorQueries(tt.args.arr, tt.args.queries))
		})
	}
}
