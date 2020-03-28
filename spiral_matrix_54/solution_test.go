package spiral_matrix_54

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "spiral order of matrix",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "spiral order of matrix",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{
				1, 2, 3, 6, 9, 8, 7, 4, 5,
			},
		},
		{
			name: "spiral order of matrix",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			want: []int{
				1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, spiralOrder(tt.args.matrix))
		})
	}
}
