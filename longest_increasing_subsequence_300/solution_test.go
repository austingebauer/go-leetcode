package longest_increasing_subsequence_300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest increasing subsequence",
			args: args{
				nums: []int{
					1, 3, 2,
				},
			},
			want: 2, // [1,2]
		},
		{
			name: "longest increasing subsequence",
			args: args{
				nums: []int{
					10, 9, 2, 5, 3, 7, 101, 18,
				},
			},
			want: 4, // [2,3,7,101]
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLIS(tt.args.nums))
		})
	}
}
