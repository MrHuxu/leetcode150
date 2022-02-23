package main

// code
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	left := (m + n + 1) / 2
	right := (m + n + 2) / 2

	return (float64(getKth(nums1, 0, m-1, nums2, 0, n-1, left)) + float64(getKth(nums1, 0, m-1, nums2, 0, n-1, right))) / 2.0
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	println(start1, end1, start2, end2, k)
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}

	if len1 == 0 {
		return nums2[start2+(k-1)]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + (min(len1, k/2) - 1)
	j := start2 + (min(len2, k/2) - 1)

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	}
	return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
