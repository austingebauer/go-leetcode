package search_in_rotated_sorted_array_33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search in rotated sorted array",
			args: args{
				nums:   []int{},
				target: 10,
			},
			want: -1,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					8, 9, 2, 3, 4,
				},
				target: 9,
			},
			want: 1,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					10,
				},
				target: 1,
			},
			want: -1,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					10,
				},
				target: 10,
			},
			want: 0,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					1, 3,
				},
				target: 4,
			},
			want: -1,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 0,
			},
			want: 4,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				nums: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.args.nums, tt.args.target))
		})
	}
}
