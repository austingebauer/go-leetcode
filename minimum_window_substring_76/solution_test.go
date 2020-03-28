package minimum_window_substring_76

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "minimum window substring",
			args: args{
				s: "ABAACBAB",
				t: "ABC",
			},
			want: "ACB",
		},
		{
			name: "minimum window substring",
			args: args{
				s: "ZADOBECOAEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "minimum window substring",
			args: args{
				s: "A",
				t: "A",
			},
			want: "A",
		},
		{
			name: "minimum window substring",
			args: args{
				s: "A",
				t: "B",
			},
			want: "",
		},
		{
			name: "minimum window substring",
			args: args{
				s: "AB",
				t: "A",
			},
			want: "A",
		},
		{
			name: "minimum window substring",
			args: args{
				s: "AB",
				t: "B",
			},
			want: "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minWindow(tt.args.s, tt.args.t))
		})
	}
}
