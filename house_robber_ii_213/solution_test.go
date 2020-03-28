package house_robber_ii_213

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "house robber ii",
			args: args{
				nums: []int{
					2, 3, 2,
				},
			},
			want: 3,
		},
		{
			name: "house robber ii",
			args: args{
				nums: []int{
					1, 2, 3, 1,
				},
			},
			want: 4,
		},
		{
			name: "house robber ii",
			args: args{
				nums: []int{
					1, 2, 1, 1,
				},
			},
			want: 3,
		},
		{
			name: "house robber ii",
			args: args{
				nums: []int{
					1, 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rob(tt.args.nums))
		})
	}
}
