package reverse_bits_190

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "reverse bits",
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
		{
			name: "reverse bits",
			args: args{
				num: 4294967293,
			},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseBits(tt.args.num))
		})
	}
}
