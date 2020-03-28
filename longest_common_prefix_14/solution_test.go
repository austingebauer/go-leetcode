package longest_common_prefix_14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "longest common prefix",
			args: args{
				strs: []string{
					"flower",
					"flow",
					"flight",
				},
			},
			want: "fl",
		},
		{
			name: "longest common prefix",
			args: args{
				strs: []string{
					"dog",
					"racecar",
					"car",
				},
			},
			want: "",
		},
		{
			name: "longest common prefix",
			args: args{
				strs: []string{
					"flower",
				},
			},
			want: "flower",
		},
		{
			name: "longest common prefix",
			args: args{
				strs: []string{
					"abcd",
					"abcd",
				},
			},
			want: "abcd",
		},
		{
			name: "longest common prefix",
			args: args{
				strs: []string{
					"a",
				},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonPrefix(tt.args.strs))
		})
	}
}
