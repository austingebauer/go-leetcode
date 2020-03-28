package construct_binary_tree_from_preorder_and_inorder_traversal_105

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "construct binary tree from preorder and inorder traversal",
			args: args{
				preorder: []int{
					3, 9, 20, 15, 7,
				},
				inorder: []int{
					9, 3, 15, 20, 7,
				},
			},
			want: &TreeNode{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, buildTree(tt.args.preorder, tt.args.inorder))
		})
	}
}
