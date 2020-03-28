package binary_tree_maximum_path_sum_124

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import . "github.com/austingebauer/go-leetcode/structures"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "binary tree maximum path sum",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 6,
		},
		{
			name: "binary tree maximum path sum",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 4,
		},
		{
			name: "binary tree maximum path sum",
			args: args{
				root: &TreeNode{
					Val: -10,
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
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxPathSum(tt.args.root))
		})
	}
}
