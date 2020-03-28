package letter_combinations_of_a_phone_number_17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "letter combinations of a phone number",
			args: args{
				digits: "23",
			},
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			name: "letter combinations of a phone number",
			args: args{
				digits: "4",
			},
			want: []string{
				"g", "h", "i",
			},
		},
		{
			name: "letter combinations of a phone number",
			args: args{
				digits: "",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, letterCombinations(tt.args.digits))
		})
	}
}
