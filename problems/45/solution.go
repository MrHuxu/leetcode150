package leetcode150

// code
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
			if dp[i+j] == 0 {
				dp[i+j] = dp[i] + 1
			} else {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
