package merge_two_sorted_lists_21

import (
	s "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists2(t *testing.T) {
	type args struct {
		l1 *s.ListNode
		l2 *s.ListNode
	}
	tests := []struct {
		name string
		args args
		want *s.ListNode
	}{
		{
			name: "merge two sorted lists",
			args: args{
				l1: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 3,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &s.ListNode{
				Val: 1,
				Next: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val: 3,
							Next: &s.ListNode{
								Val: 4,
								Next: &s.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists2(tt.args.l1, tt.args.l2))
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *s.ListNode
		l2 *s.ListNode
	}
	tests := []struct {
		name string
		args args
		want *s.ListNode
	}{
		{
			name: "merge two sorted lists",
			args: args{
				l1: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 3,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &s.ListNode{
				Val: 1,
				Next: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val: 3,
							Next: &s.ListNode{
								Val: 4,
								Next: &s.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists0(tt.args.l1, tt.args.l2))
		})
	}
}
