package leetcode150

// code
var count = []int{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	var result []byte

	used := make([]bool, n+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			if !used[j] {
				if k == 1 {
					result = append(result, byte(j+48))
					used[j] = true
					break
				} else {
					if k <= count[n-1-i] {
						result = append(result, byte(j+48))
						used[j] = true
						break
					} else {
						k -= count[n-i-1]
					}
				}
			}
		}
	}

	return string(result)
}
