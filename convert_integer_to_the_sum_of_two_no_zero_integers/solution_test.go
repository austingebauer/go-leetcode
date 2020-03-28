package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getNoZeroIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "convert integer to sum of two non-zero integers",
			args: args{
				n: 2,
			},
			want: []int{
				1, 1,
			},
		},
		{
			name: "convert integer to sum of two non-zero integers",
			args: args{
				n: 11,
			},
			want: []int{
				2, 9,
			},
		},
		{
			name: "convert integer to sum of two non-zero integers",
			args: args{
				n: 10000,
			},
			want: []int{
				1, 9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getNoZeroIntegers(tt.args.n))
		})
	}
}
