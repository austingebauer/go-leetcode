package add_two_numbers_2

import (
	"github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *structures.ListNode
		l2 *structures.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structures.ListNode
	}{
		{
			name: "add two numbers",
			args: args{
				l1: &structures.ListNode{
					Val: 2,
					Next: &structures.ListNode{
						Val: 4,
						Next: &structures.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &structures.ListNode{
					Val: 5,
					Next: &structures.ListNode{
						Val: 6,
						Next: &structures.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &structures.ListNode{
				Val: 7,
				Next: &structures.ListNode{
					Val: 0,
					Next: &structures.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "add two numbers",
			args: args{
				l1: &structures.ListNode{
					Val: 1,
					Next: &structures.ListNode{
						Val: 0,
						Next: &structures.ListNode{
							Val: 0,
							Next: &structures.ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
				l2: &structures.ListNode{
					Val: 5,
					Next: &structures.ListNode{
						Val: 6,
						Next: &structures.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &structures.ListNode{
				Val: 6,
				Next: &structures.ListNode{
					Val: 6,
					Next: &structures.ListNode{
						Val: 4,
						Next: &structures.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(tt.args.l1, tt.args.l2))
			assert.Equal(t, tt.want, addTwoNumbers2(tt.args.l1, tt.args.l2))
		})
	}
}
