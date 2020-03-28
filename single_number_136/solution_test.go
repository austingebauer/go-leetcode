package single_number_136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single number in list",
			args: args{
				nums: []int{
					2, 2, 1,
				},
			},
			want: 1,
		},
		{
			name: "single number in list",
			args: args{
				nums: []int{
					2, 1, 2,
				},
			},
			want: 1,
		},
		{
			name: "single number in list",
			args: args{
				nums: []int{
					4, 1, 2, 1, 2,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, singleNumber(tt.args.nums))
			assert.Equal(t, tt.want, singleNumber2(tt.args.nums))
		})
	}
}
