package number_of_1_bits_191

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hamming weight",
			args: args{
				num: 11,
			},
			want: 3,
		},
		{
			name: "hamming weight",
			args: args{
				num: 128,
			},
			want: 1,
		},
		{
			name: "hamming weight",
			args: args{
				num: 4294967293,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hammingWeight(tt.args.num))
		})
	}
}
