package non_overlapping_intervals_435

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "non-overlapping intervals",
			args: args{
				intervals: [][]int{
					{1, 2}, {2, 3}, {3, 4}, {1, 3},
				},
			},
			want: 1,
		},
		{
			name: "non-overlapping intervals",
			args: args{
				intervals: [][]int{
					{1, 2}, {1, 2}, {1, 2},
				},
			},
			want: 2,
		},
		{
			name: "non-overlapping intervals",
			args: args{
				intervals: [][]int{
					{1, 2}, {2, 3},
				},
			},
			want: 0,
		},
		{
			name: "non-overlapping intervals",
			args: args{
				intervals: [][]int{
					{1, 100}, {11, 22}, {2, 12}, {1, 11},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, eraseOverlapIntervals(tt.args.intervals))
		})
	}
}
