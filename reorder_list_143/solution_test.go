package reorder_list_143

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "reorder list",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
		{
			name: "reorder list",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			assert.Equal(t, tt.want, tt.args.head)
		})
	}
}
