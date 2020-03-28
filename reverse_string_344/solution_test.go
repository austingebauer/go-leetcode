package reverse_string_344

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "reverse string",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
		{
			name: "reverse string",
			args: args{
				s: []byte("Hannah"),
			},
			want: []byte("hannaH"),
		},
		{
			name: "reverse string",
			args: args{
				s: []byte("a"),
			},
			want: []byte("a"),
		},
		{
			name: "reverse string",
			args: args{
				s: []byte(""),
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !assert.Equal(t, tt.want, tt.args.s) {
				fmt.Printf("tt.want: %v \n", tt.want)
				fmt.Printf("tt.got:  %v \n", tt.args.s)
			}
		})
	}
}
