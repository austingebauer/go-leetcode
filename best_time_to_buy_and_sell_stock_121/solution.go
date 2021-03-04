package best_time_to_buy_and_sell_stock_121

import "math"

// O(n) time and O(1) space
func maxProfit3(prices []int) int {
	var res float64
	var low = math.MaxFloat64
	for _, p := range prices {
		res = math.Max(res, float64(p)-low)
		low = math.Min(low, float64(p))
	}
	return int(res)
}

// O(n) time and O(1) space
func maxProfit2(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// O(n^2) time and O(1) space
func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		buyPrice := prices[i]
		maxSellPrice := 0

		for j := i + 1; j < len(prices); j++ {
			sellPrice := prices[j]

			if sellPrice > buyPrice {
				maxSellPrice = int(math.Max(float64(maxSellPrice), float64(sellPrice)))
			}
		}

		maxProfit = int(math.Max(float64(maxProfit), float64(maxSellPrice-buyPrice)))
	}

	return maxProfit
}
