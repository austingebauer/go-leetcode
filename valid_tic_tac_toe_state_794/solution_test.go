package valid_tic_tac_toe_state_794

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"O  ", "   ", "   ",
				},
			},
			want: false,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XOX", " X ", "   ",
				},
			},
			want: false,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XXX", "   ", "OOO",
				},
			},
			want: false,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XXX", " O ", "  O",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"OX ", " OX", " XO",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XOX", "O O", "XOX",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"   ", "   ", "   ",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XOX", "XXO", "OXO",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XOX", "XXO", "OOX",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XXX", "OOX", "OOX",
				},
			},
			want: true,
		},
		{
			name: "valid tic-tac-toe state",
			args: args{
				board: []string{
					"XXX", "XOO", "OO ",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validTicTacToe(tt.args.board))
		})
	}
}
