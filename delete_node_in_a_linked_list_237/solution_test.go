package delete_node_in_a_linked_list_237

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node     *ListNode
		toDelete int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "delete node in linked list",
			args: args{
				node: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
				toDelete: 5,
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "delete node in linked list",
			args: args{
				node: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
				toDelete: 1,
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// advance a pointer to the node we should point at for the test
			c := tt.args.node
			for c.Val != tt.args.toDelete {
				c = c.Next
			}

			deleteNode(c)
			assert.Equal(t, tt.want, tt.args.node)
		})
	}
}
