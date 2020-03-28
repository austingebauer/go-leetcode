package insert_interval_57

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 9},
				},
				newInterval: []int{
					5, 6,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 6},
				{7, 9},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 9},
				},
				newInterval: []int{
					5, 7,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 9},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 9},
				},
				newInterval: []int{
					5, 9,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 9},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 8},
				},
				newInterval: []int{
					5, 9,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 9},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 8},
					{11, 12},
				},
				newInterval: []int{
					5, 10,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 10},
				{11, 12},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{2, 4},
					{7, 8},
					{11, 12},
				},
				newInterval: []int{
					5, 11,
				},
			},
			want: [][]int{
				{2, 4},
				{5, 12},
			},
		},
		{
			name: "insert interval",
			args: args{
				intervals: [][]int{
					{1, 5},
				},
				newInterval: []int{
					6, 8,
				},
			},
			want: [][]int{
				{1, 5},
				{6, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, insert(tt.args.intervals, tt.args.newInterval))
		})
	}
}
