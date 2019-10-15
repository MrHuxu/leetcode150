package leetcode150

// code
func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b))+1)

	var plus byte
	for i := 0; i < len(a) || i < len(b); i++ {
		sum := plus

		if i < len(a) {
			sum += a[len(a)-1-i] - 48
		}
		if i < len(b) {
			sum += b[len(b)-1-i] - 48
		}

		result[len(result)-1-i] = (sum & 1) + 48
		plus = sum >> 1
	}
	if plus == 1 {
		result[0] = '1'
	} else {
		result = result[1:len(result)]
	}

	return string(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
