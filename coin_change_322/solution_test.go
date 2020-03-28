package coin_change_322

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coin change",
			args: args{
				coins: []int{
					1, 2,
				},
				amount: 3,
			},
			want: 2,
		},
		{
			name: "coin change",
			args: args{
				coins: []int{
					1, 2, 5,
				},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "coin change",
			args: args{
				coins: []int{
					1, 2,
				},
				amount: 3,
			},
			want: 2,
		},
		{
			name: "coin change",
			args: args{
				coins: []int{
					3,
				},
				amount: 1,
			},
			want: -1,
		},
		{
			name: "coin change",
			args: args{
				coins: []int{
					186, 419, 83, 408,
				},
				amount: 6249,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, coinChange(tt.args.coins, tt.args.amount))
		})
	}
}
