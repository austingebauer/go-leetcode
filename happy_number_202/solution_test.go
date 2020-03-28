package happy_number_202

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is happy number",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "is happy number",
			args: args{
				n: 7,
			},
			want: true,
		},
		{
			name: "is happy number",
			args: args{
				n: 116,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isHappy(tt.args.n))
		})
	}
}
