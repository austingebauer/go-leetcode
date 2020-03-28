package maximum_depth_of_binary_tree_104

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum depth of binary tree",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "maximum depth of binary tree",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: -1,
							},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "maximum depth of binary tree",
			args: args{
				root: &TreeNode{
					Val: 3,
				},
			},
			want: 1,
		},
		{
			name: "maximum depth of binary tree",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxDepth(tt.args.root))
		})
	}
}
