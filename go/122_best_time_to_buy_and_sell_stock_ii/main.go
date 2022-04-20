package main

// code
func maxProfit(prices []int) int {
	var profit int

	for i := 1; i < len(prices); i++ {
		if sub := prices[i] - prices[i-1]; sub > 0 {
			profit += sub
		}
	}

	return profit
}
