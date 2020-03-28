package container_with_most_water_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "container with most water",
			args: args{
				height: []int{
					1, 8, 6, 2, 5, 4, 8, 3, 7,
				},
			},
			want: 49,
		},
		{
			name: "container with most water",
			args: args{
				height: []int{
					1, 1,
				},
			},
			want: 1,
		},
		{
			name: "container with most water",
			args: args{
				height: []int{
					1, 2, 4, 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.height))
		})
	}
}
