package group_anagrams_49

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "group anagrams",
			args: args{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			want: [][]string{
				{
					"bat",
				},
				{
					"eat", "tea", "ate",
				},
				{
					"tan", "nat",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := groupAnagrams(tt.args.strs)
			sort.Slice(g, func(i int, j int) bool {
				return g[i][0] < g[j][0]
			})
			assert.Equal(t, tt.want, g)
		})
	}
}
