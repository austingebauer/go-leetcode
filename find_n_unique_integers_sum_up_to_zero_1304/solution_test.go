package find_n_unique_integers_sum_up_to_zero_1304

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find n unique integers that sum up to zero",
			args: args{
				n: 5,
			},
			want: []int{
				1, -1, 2, -2, 0,
			},
		},
		{
			name: "find n unique integers that sum up to zero",
			args: args{
				n: 3,
			},
			want: []int{
				1, -1, 0,
			},
		},
		{
			name: "find n unique integers that sum up to zero",
			args: args{
				n: 1,
			},
			want: []int{
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumZero(tt.args.n))
			assert.Equal(t, tt.want, sumZeroNoAppend(tt.args.n))
		})
	}
}
