package kth_smallest_element_in_a_bst_230

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Kth Smallest Element in a BST",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "Kth Smallest Element in a BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				k: 3,
			},
			want: 3,
		},
		{
			name: "Kth Smallest Element in a BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				k: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, kthSmallest(tt.args.root, tt.args.k))
		})
	}
}
