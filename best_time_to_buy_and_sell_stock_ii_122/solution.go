package best_time_to_buy_and_sell_stock_ii_122

// Keep on adding the difference between the consecutive numbers of the slice,
// if the second number is larger than the first one (profit). The image in the
// readme helps to visualize why this works. The problem description could be
// improved, as it doesn't say whether or not selling and buying back on the
// same day is possible. For example, given [2, 5, 8], are we allowed to buy
// at 2, sell at 5, buy back at 5, sell at 8? We can, but the description
// doesn't make that clear. O(n) time and O(1) space.
func maxProfit(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
