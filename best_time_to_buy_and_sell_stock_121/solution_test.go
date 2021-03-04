package best_time_to_buy_and_sell_stock_121

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
			name: "best time to buy and sell stock",
			args: args{
				prices: []int{
					7, 1, 5, 3, 6, 4,
				},
			},
			want: 5,
		},
		{
			name: "best time to buy and sell stock",
			args: args{
				prices: []int{
					7, 6, 4, 3, 1,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProfit3(tt.args.prices))
			assert.Equal(t, tt.want, maxProfit2(tt.args.prices))
			assert.Equal(t, tt.want, maxProfit(tt.args.prices))
		})
	}
}
