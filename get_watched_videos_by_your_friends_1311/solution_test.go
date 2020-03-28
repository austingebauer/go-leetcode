package get_watched_videos_by_your_friends_1311

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_watchedVideosByFriends(t *testing.T) {
	type args struct {
		watchedVideos [][]string
		friends       [][]int
		id            int
		level         int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				watchedVideos: [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
				friends:       [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
				id:            0,
				level:         1,
			},
			want: []string{"B", "C"},
		},
		{
			name: "",
			args: args{
				watchedVideos: [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
				friends:       [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
				id:            0,
				level:         2,
			},
			want: []string{"D"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, watchedVideosByFriends(tt.args.watchedVideos,
				tt.args.friends, tt.args.id, tt.args.level))
		})
	}
}
