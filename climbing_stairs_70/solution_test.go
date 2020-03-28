package climbing_stairs_70

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "climbing stairs",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "climbing stairs",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "climbing stairs",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "climbing stairs",
			args: args{
				n: 44,
			},
			want: 1134903170,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, climbStairs(tt.args.n))
		})
	}
}
