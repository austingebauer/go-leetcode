package reverse_string_ii_541

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "",
			args: args{
				s: "abcdefgi",
				k: 2,
			},
			want: "bacdfegi",
		},
		{
			name: "",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
		{
			name: "",
			args: args{
				s: "abcd",
				k: 4,
			},
			want: "dcba",
		},
		{
			name: "",
			args: args{
				s: "abcdefg",
				k: 3,
			},
			want: "cbadefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
