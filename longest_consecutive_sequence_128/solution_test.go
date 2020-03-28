package longest_consecutive_sequence_128

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest consecutive sequence",
			args: args{
				nums: []int{
					100, 4, 200, 1, 3, 2,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestConsecutive(tt.args.nums))
		})
	}
}
