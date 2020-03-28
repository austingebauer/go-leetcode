package reverse_linked_list_206

import (
	s "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *s.ListNode
	}
	tests := []struct {
		name string
		args args
		want *s.ListNode
	}{
		{
			name: "reverse linked list",
			args: args{
				head: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val: 3,
							Next: &s.ListNode{
								Val: 4,
								Next: &s.ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &s.ListNode{
				Val: 5,
				Next: &s.ListNode{
					Val: 4,
					Next: &s.ListNode{
						Val: 3,
						Next: &s.ListNode{
							Val: 2,
							Next: &s.ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "reverse linked list",
			args: args{
				head: &s.ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: &s.ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "reverse linked list",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "reverse linked list",
			args: args{
				head: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &s.ListNode{
				Val: 2,
				Next: &s.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(tt.args.head))
		})
	}
}
