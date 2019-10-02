package leetcode150

// code
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	totalNums := make([]int, totalLen)

	var i, j int
	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) {
			totalNums[i+j] = nums2[j]
			j++
			continue
		}

		if j == len(nums2) {
			totalNums[i+j] = nums1[i]
			i++
			continue
		}

		if nums1[i] <= nums2[j] {
			totalNums[i+j] = nums1[i]
			i++
		} else {
			totalNums[i+j] = nums2[j]
			j++
		}

		if i+j > totalLen/2 {
			break
		}
	}
	if totalLen%2 == 0 {
		return (float64(totalNums[totalLen/2-1]) + float64(totalNums[totalLen/2])) / 2.0
	}
	return float64(totalNums[totalLen/2])
}
