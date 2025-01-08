package flatten_nested_list_iterator_341

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	type args struct {
		list []*NestedInteger
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				// [[1,1],2,[1,1]]
				list: []*NestedInteger{
					NewNestedInteger([]*NestedInteger{
						NewNestedInteger(1),
						NewNestedInteger(1),
					}),
					NewNestedInteger(2),
					NewNestedInteger([]*NestedInteger{
						NewNestedInteger(1),
						NewNestedInteger(1),
					}),
				},
			},
			want: []int{1, 1, 2, 1, 1},
		},
		{
			name: "1",
			args: args{
				// [1,[4,[6]]]
				list: []*NestedInteger{
					NewNestedInteger(1),
					NewNestedInteger([]*NestedInteger{
						NewNestedInteger(4),
						NewNestedInteger([]*NestedInteger{
							NewNestedInteger(6),
						}),
					}),
				},
			},
			want: []int{1, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int

			itr := Constructor(tt.args.list)
			for itr.HasNext() {
				got = append(got, itr.Next())
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
