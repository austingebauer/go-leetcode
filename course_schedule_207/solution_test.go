package course_schedule_207

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "course schedule",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{0, 1},
					{1, 2},
					{2, 0},
				},
			},
			want: false,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{0, 1},
					{0, 2},
					{1, 2},
				},
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{0, 2},
					{1, 0},
				},
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{0, 2},
				},
			},
			want: true,
		},
		{
			name: "course schedule",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{0, 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canFinishDFS(tt.args.numCourses, tt.args.prerequisites))
			assert.Equal(t, tt.want, canFinishBFS(tt.args.numCourses, tt.args.prerequisites))
		})
	}
}
