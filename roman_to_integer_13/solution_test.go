package roman_to_integer_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "roman numeral to int",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "roman numeral to int",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "roman numeral to int",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "roman numeral to int",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "roman numeral to int",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, romanToInt(tt.args.s))
		})
	}
}
