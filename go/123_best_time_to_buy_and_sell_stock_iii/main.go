package main

import "math"

// code
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy1, sell1, buy2, sell2 := math.MinInt32, 0, math.MinInt32, 0
	for _, price := range prices {
		buy1 = max(buy1, -price)
		sell1 = max(sell1, buy1+price)
		buy2 = max(buy2, sell1-price)
		sell2 = max(sell2, buy2+price)
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
