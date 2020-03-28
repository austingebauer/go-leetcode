package counting_bits_338

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "counting bits",
			args: args{
				num: 2,
			},
			want: []int{
				0, 1, 1,
			},
		},
		{
			name: "counting bits",
			args: args{
				num: 5,
			},
			want: []int{
				0, 1, 1, 2, 1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBits(tt.args.num))
		})
	}
}
