package sum_of_two_integers_371

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of two integers",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "sum of two integers",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "sum of two integers",
			args: args{
				a: -2,
				b: 3,
			},
			want: 1,
		},
		{
			name: "sum of two integers",
			args: args{
				a: -2,
				b: -3,
			},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getSum(tt.args.a, tt.args.b))
		})
	}
}
