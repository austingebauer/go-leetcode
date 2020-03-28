package invert_binary_tree_226

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "invert binary tree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "invert binary tree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			name: "invert binary tree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			name: "invert binary tree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, invertTree(tt.args.root))
		})
	}
}
