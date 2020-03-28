package set_matrix_zeros_73

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "set matrix zeros",
			args: args{
				matrix: [][]int{
					{1, 1, 0},
					{1, 1, 2},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 1, 0},
			},
		},
		{
			name: "set matrix zeros",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{0, 1, 2},
				},
			},
			want: [][]int{
				{0, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "set matrix zeros",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "set matrix zeros",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			assert.Equal(t, tt.want, tt.args.matrix)
		})
	}
}
