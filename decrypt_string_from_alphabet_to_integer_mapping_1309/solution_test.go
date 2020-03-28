package decrypt_string_from_alphabet_to_integer_mapping_1309

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_freqAlphabets(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "freqAlphabets",
			args: args{
				s: "10#11#12",
			},
			want: "jkab",
		},
		{
			name: "freqAlphabets",
			args: args{
				s: "1326#",
			},
			want: "acz",
		},
		{
			name: "freqAlphabets",
			args: args{
				s: "25#",
			},
			want: "y",
		},
		{
			name: "freqAlphabets",
			args: args{
				s: "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, freqAlphabets(tt.args.s))
		})
	}
}
