package find_the_difference_of_two_arrays_2215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {4, 6}},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			want: [][]int{{3}, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)
			assert.ElementsMatch(t, tt.want[0], got[0])
			assert.ElementsMatch(t, tt.want[1], got[1])
		})
	}
}
