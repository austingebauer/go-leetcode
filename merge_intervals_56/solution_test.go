package merge_intervals_56

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "merge intervals 1",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "merge intervals 2",
			args: args{
				intervals: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "merge intervals 3",
			args: args{
				intervals: [][]int{
					{1, 4},
					{0, 4},
				},
			},
			want: [][]int{
				{0, 4},
			},
		},
		{
			name: "merge intervals 4",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 3},
				},
			},
			want: [][]int{
				{1, 4},
			},
		},
		{
			name: "merge intervals 5",
			args: args{
				intervals: [][]int{
					{1, 4},
					{0, 2},
					{3, 5},
				},
			},
			want: [][]int{
				{0, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, merge(tt.args.intervals))
		})
	}
}
