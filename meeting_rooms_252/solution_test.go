package meeting_rooms_252

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "meeting rooms",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: false,
		},
		{
			name: "meeting rooms",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canAttendMeetings(tt.args.intervals))
		})
	}
}
