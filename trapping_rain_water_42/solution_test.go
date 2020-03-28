package trapping_rain_water_42

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trapping rain water",
			args: args{
				height: []int{
					0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, trap(tt.args.height))
			assert.Equal(t, tt.want, trap2(tt.args.height))
		})
	}
}
