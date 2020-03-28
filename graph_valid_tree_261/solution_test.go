package graph_valid_tree_261

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "graph valid tree",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1},
					{0, 2},
					{0, 3},
					{1, 4},
				},
			},
			want: true,
		},
		{
			name: "graph valid tree",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{2, 3},
					{1, 3},
					{1, 4},
				},
			},
			want: false,
		},
		{
			name: "graph valid tree",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1},
					{0, 4},
					{1, 4},
					{2, 3},
				},
			},
			want: false,
		},
		{
			name: "graph valid tree",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validTree(tt.args.n, tt.args.edges))
		})
	}
}
