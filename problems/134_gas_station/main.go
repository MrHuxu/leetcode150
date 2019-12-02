package main

// code
func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, totalCost int

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalCost > totalGas {
		return -1
	}

	var result int

	var gasSum, costSum int
	for i := 0; i < len(gas); i++ {
		gasSum += gas[i]
		costSum += cost[i]
		if costSum > gasSum {
			gasSum = 0
			costSum = 0
			result = i + 1
		}
	}

	return result
}
