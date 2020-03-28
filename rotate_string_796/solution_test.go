package rotate_string_796

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				A: "bbbacddceeb",
				B: "ceebbbbacdd",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rotateString(tt.args.A, tt.args.B))
		})
	}
}
