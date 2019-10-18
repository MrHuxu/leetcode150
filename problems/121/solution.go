package leetcode150

// code
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var profit int

	min := prices[0]
	for _, price := range prices {
		if price > min {
			profit = max(profit, price-min)
		} else {
			min = price
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
