package product_of_array_except_self_238

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "product of array except self",
			args: args{
				nums: []int{
					1, 2, 3, 4,
				},
			},
			want: []int{
				24, 12, 8, 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, productExceptSelf(tt.args.nums))
		})
	}
}
