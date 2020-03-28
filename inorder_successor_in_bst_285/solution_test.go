package inorder_successor_in_bst_285

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "in-order successor of BST",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				p: &TreeNode{
					Val: 1,
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "in-order successor of BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Left: &TreeNode{
								Val: 1,
							},
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				p: &TreeNode{
					Val: 6,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, inorderSuccessor(tt.args.root, tt.args.p))
		})
	}
}
