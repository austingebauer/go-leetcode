package word_search_79

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A'},
				},
				word: "A",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A'},
					{'A'},
				},
				word: "A",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A'},
					{'A'},
				},
				word: "AA",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'A'},
					{'A'},
				},
				word: "AB",
			},
			want: false,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, exist(tt.args.board, tt.args.word))
		})
	}
}
