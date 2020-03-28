package validate_binary_search_tree_98

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validate binary search tree",
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
			},
			want: true,
		},
		{
			name: "validate binary search tree",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: false,
		},
		// [10,5,15,null,null,6,20]
		{
			name: "validate binary search tree",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 20,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidBST(tt.args.root))
		})
	}
}
