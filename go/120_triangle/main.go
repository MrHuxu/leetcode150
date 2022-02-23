package main

// code
func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			switch {
			case j == 0:
				triangle[i][j] += triangle[i-1][j]

			case j == len(triangle[i])-1:
				triangle[i][j] += triangle[i-1][j-1]

			case j > 0 && j < len(triangle[i])-1:
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}

	return min(triangle[len(triangle)-1]...)
}

func min(nums ...int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < result {
			result = nums[i]
		}
	}

	return result
}
