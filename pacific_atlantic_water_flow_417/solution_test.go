package pacific_atlantic_water_flow_417

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "pacific atlantic water flow",
			args: args{
				matrix: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "pacific atlantic water flow",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
			},
			want: [][]int{
				{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pacificAtlantic(tt.args.matrix))
			assert.Equal(t, tt.want, pacificAtlantic0(tt.args.matrix))
		})
	}
}
