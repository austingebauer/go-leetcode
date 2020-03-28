package encode_and_decode_strings_271

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "encode and decode strings",
			args: args{
				strs: []string{
					"should", "encode", "and", "",
				},
			},
			want: []string{
				"should", "encode", "and", "",
			},
		},
		{
			name: "encode and decode strings",
			args: args{
				strs: []string{
					"63/Rc", "h", "BmI3FS~J9#vmk", "7uBZ?7*/", "24h+X", "O ",
				},
			},
			want: []string{
				"63/Rc", "h", "BmI3FS~J9#vmk", "7uBZ?7*/", "24h+X", "O ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Codec
			e := c.Encode(tt.args.strs)
			d := c.Decode(e)
			assert.Equal(t, tt.want, d)
			assert.Equal(t, d, tt.args.strs)
		})
	}
}
