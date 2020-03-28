package symmetric_tree_101

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "symmetric tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "symmetric tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSymmetric(tt.args.root))
		})
	}
}
