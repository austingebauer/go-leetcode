package critical_connections_in_a_network_1192

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_criticalConnections(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "critical connections in a network",
			args: args{
				n: 4,
				connections: [][]int{
					{0, 1}, {1, 2}, {2, 0}, {1, 3},
				},
			},
			want: [][]int{
				{1, 3}, // [[3,1]] also accepted
			},
		},
		{
			name: "critical connections in a network",
			args: args{
				n: 6,
				connections: [][]int{
					{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}, {4, 5}, {5, 3},
				},
			},
			want: [][]int{
				{1, 3}, // [[3,1]] also accepted
			},
		},
		{
			name: "critical connections in a network",
			args: args{
				n: 5,
				connections: [][]int{
					{1, 0}, {2, 0}, {3, 2}, {4, 2}, {4, 3}, {3, 0}, {4, 0},
				},
			},
			want: [][]int{
				{1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, criticalConnections(tt.args.n, tt.args.connections))
		})
	}
}
