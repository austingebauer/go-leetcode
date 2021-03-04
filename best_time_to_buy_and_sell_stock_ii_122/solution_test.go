package best_time_to_buy_and_sell_stock_ii_122

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "best time to buy and sell stock ii",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "best time to buy and sell stock ii",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "best time to buy and sell stock ii",
			args: args{
				prices: []int{2, 7, 1, 8},
			},
			want: 12,
		},
		{
			name: "best time to buy and sell stock ii",
			args: args{
				prices: []int{2, 5, 8},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProfit(tt.args.prices))
		})
	}
}
