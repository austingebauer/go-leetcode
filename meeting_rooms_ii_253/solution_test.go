package meeting_rooms_ii_253

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "meeting rooms ii 1",
			args: args{
				intervals: [][]int{
					{0, 30},
				},
			},
			want: 1,
		},
		{
			name: "meeting rooms ii 2",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "meeting rooms ii 3",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
		{
			name: "meeting rooms ii 4",
			args: args{
				intervals: [][]int{
					{9, 10},
					{4, 9},
					{4, 17},
				},
			},
			want: 2,
		},
		{
			name: "meeting rooms ii 5",
			args: args{
				intervals: [][]int{
					{1, 5},
					{8, 9},
					{8, 9},
				},
			},
			want: 2,
		},
		{
			name: "meeting rooms ii 6",
			args: args{
				intervals: [][]int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minMeetingRooms(tt.args.intervals))
			assert.Equal(t, tt.want, minMeetingRooms0(tt.args.intervals))
		})
	}
}
