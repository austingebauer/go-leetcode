package longest_repeating_character_replacement_424

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest repeating character replacement",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
		{
			name: "longest repeating character replacement",
			args: args{
				s: "BAAAB",
				k: 2,
			},
			want: 5,
		},
		{
			name: "longest repeating character replacement",
			args: args{
				s: "ABBB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "longest repeating character replacement",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, characterReplacement(tt.args.s, tt.args.k))
		})
	}
}
