package merge_sorted_array_88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{
					1, 2, 3, 0, 0, 0,
				},
				m: 3,
				nums2: []int{
					2, 5, 6,
				},
				n: 3,
			},
			want: []int{
				1, 2, 2, 3, 5, 6,
			},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{
					1, 0,
				},
				m: 1,
				nums2: []int{
					16,
				},
				n: 1,
			},
			want: []int{
				1, 16,
			},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{
					16, 0,
				},
				m: 1,
				nums2: []int{
					1,
				},
				n: 1,
			},
			want: []int{
				1, 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//merge0(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			assert.Equal(t, tt.want, tt.args.nums1)
		})
	}
}
