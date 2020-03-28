package count_and_say_38

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "count and say",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "count and say",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "count and say",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "count and say",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "count and say",
			args: args{
				n: 5,
			},
			want: "111221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countAndSay(tt.args.n))
		})
	}
}
