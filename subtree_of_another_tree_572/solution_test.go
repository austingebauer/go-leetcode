package subtree_of_another_tree_572

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is subtree",
			args: args{
				s: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
					Val: 3,
				},
				t: &TreeNode{
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
					Val: 4,
				},
			},
			want: true,
		},
		{
			name: "is not subtree",
			args: args{
				s: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 0,
							},
						},
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
					Val: 3,
				},
				t: &TreeNode{
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
					Val: 4,
				},
			},
			want: false,
		},
		{
			name: "is subtree with duplicate values in tree",
			args: args{
				s: &TreeNode{
					Left: &TreeNode{
						Val: 1,
					},
					Val: 1,
				},
				t: &TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSubtree(tt.args.s, tt.args.t))
		})
	}
}
