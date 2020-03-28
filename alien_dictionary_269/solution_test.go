package alien_dictionary_269

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alien dictionary",
			args: args{
				words: []string{
					"wrt",
					"wrf",
					"er",
					"ett",
					"rftt",
				},
			},
			want: "wertf",
		},
		{
			name: "alien dictionary",
			args: args{
				words: []string{
					"z",
					"x",
				},
			},
			want: "zx",
		},
		{
			name: "alien dictionary",
			args: args{
				words: []string{
					"z",
					"x",
					"z",
				},
			},
			want: "",
		},
		{
			name: "alien dictionary",
			args: args{
				words: []string{
					"za",
					"zb",
					"ca",
					"cb",
				},
			},
			want: "azbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, alienOrder(tt.args.words))
		})
	}
}
