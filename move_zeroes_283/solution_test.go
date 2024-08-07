package move_zeroes_283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "move zeros",
			args: args{
				nums: []int{
					0, 1, 0, 3, 12,
				},
			},
			want: []int{
				1, 3, 12, 0, 0,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					12, 0, 0,
				},
			},
			want: []int{
				12, 0, 0,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					0, 1, 0,
				},
			},
			want: []int{
				1, 0, 0,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					0, 0, 2,
				},
			},
			want: []int{
				2, 0, 0,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					1, 0, 2,
				},
			},
			want: []int{
				1, 2, 0,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					1, 1, 1,
				},
			},
			want: []int{
				1, 1, 1,
			},
		},
		{
			name: "move zeros",
			args: args{
				nums: []int{
					0, 0, 0,
				},
			},
			want: []int{
				0, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
