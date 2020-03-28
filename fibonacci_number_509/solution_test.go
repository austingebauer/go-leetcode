package fibonacci_number_509

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nth fibonacci number",
			args: args{
				N: 0,
			},
			want: 0,
		},
		{
			name: "nth fibonacci number",
			args: args{
				N: 1,
			},
			want: 1,
		},
		{
			name: "nth fibonacci number",
			args: args{
				N: 2,
			},
			want: 1,
		},
		{
			name: "nth fibonacci number",
			args: args{
				N: 3,
			},
			want: 2,
		},
		{
			name: "nth fibonacci number",
			args: args{
				N: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fib(tt.args.N))
		})
	}
}
