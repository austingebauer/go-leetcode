package lowest_common_ancestor_of_a_binary_search_tree_235

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lowest common ancestor",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: 2,
				q: 8,
			},
			want: 6,
		},
		{
			name: "lowest common ancestor",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: 2,
				q: 4,
			},
			want: 2,
		},
		{
			name: "lowest common ancestor",
			args: args{
				root: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 8,
						Right: &TreeNode{
							Val: 10,
							Right: &TreeNode{
								Val: 12,
							},
							Left: &TreeNode{
								Val: 9,
							},
						},
					},
				},
				p: 8,
				q: 9,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := findNodeBST(tt.args.root, tt.args.p)
			q := findNodeBST(tt.args.root, tt.args.q)
			assert.Equal(t, tt.want, lowestCommonAncestor(tt.args.root, p, q).Val)
		})
	}
}

// Assumes the passed tree is a BST, has all unique values, and v is in the tree
func findNodeBST(root *TreeNode, v int) *TreeNode {
	curr := root
	for curr.Val != v {
		if v < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return curr
}
