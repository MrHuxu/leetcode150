package main

// code
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m
	j := n
	reverseIndex := 1

	for i > 0 && j > 0 {
		if nums1[i-1] >= nums2[j-1] {
			nums1[m+n-reverseIndex] = nums1[i-1]
			i--
		} else {
			nums1[m+n-reverseIndex] = nums2[j-1]
			j--
		}

		reverseIndex++
	}

	for j > 0 {
		nums1[m+n-reverseIndex] = nums2[j-1]
		j--
		reverseIndex++
	}
}
