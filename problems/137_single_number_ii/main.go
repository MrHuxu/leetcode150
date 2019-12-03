package main

// code
func singleNumber(nums []int) int {
	var result int32

	for i := uint(0); i < 32; i++ {
		var cnt int

		for _, num := range nums {
			cnt += (num >> i) & 1
		}

		if cnt%3 != 0 {
			result = result | (1 << i)
		}
	}

	return int(result)
}
