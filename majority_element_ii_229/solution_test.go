package majority_element_ii_229

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "majority element ii",
			args: args{
				nums: []int{
					3, 2, 3,
				},
			},
			want: []int{
				3,
			},
		},
		{
			name: "majority element ii",
			args: args{
				nums: []int{
					1, 1, 1, 3, 3, 2, 2, 2,
				},
			},
			want: []int{
				1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, majorityElement(tt.args.nums))
		})
	}
}
