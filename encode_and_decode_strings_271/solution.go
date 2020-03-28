package encode_and_decode_strings_271

import (
	"bytes"
	"fmt"
	"strconv"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var b bytes.Buffer
	for _, s := range strs {
		// encode string using length and underscore header
		l := len(s)
		b.WriteString(fmt.Sprintf("%d_%s", l, s))
	}
	return b.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	s := make([]string, 0)
	for i := 0; i < len(strs); i++ {
		var b bytes.Buffer

		// read string header until an underscore is seen
		charCnt := 0
		for string(strs[i]) != "_" {
			l, _ := strconv.Atoi(string(strs[i]))
			charCnt = (charCnt + l) * 10
			i++
		}
		charCnt /= 10

		// add charCnt characters to the buffer
		for j := 1; j <= charCnt; j++ {
			b.WriteByte(strs[i+j])
		}
		s = append(s, b.String())

		// move i forward over the last read string
		i += charCnt
	}

	return s
}
