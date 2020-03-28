package majority_element_169

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
		want int
	}{
		{
			name: "majority element",
			args: args{
				[]int{
					3, 3,
				},
			},
			want: 3,
		},
		{
			name: "majority element",
			args: args{
				[]int{
					3, 2, 3,
				},
			},
			want: 3,
		},
		{
			name: "majority element",
			args: args{
				[]int{
					2, 2, 1, 1, 1, 2, 2,
				},
			},
			want: 2,
		},
		{
			name: "majority element",
			args: args{
				[]int{
					7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7,
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, majorityElement(tt.args.nums))
		})
	}
}
