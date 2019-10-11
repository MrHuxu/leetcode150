package leetcode150

// code
func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	tmp := []byte{0}
	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			currPos := (len(num2) - 1 - i) + (len(num1) - 1 - j)
			nextPos := (len(num2) - 1 - i) + (len(num1) - 1 - j) + 1

			product := (num1[j] - 48) * (num2[i] - 48)

			if currPos > len(tmp)-1 {
				tmp = append(tmp, product)
			} else {
				tmp[currPos] += product
			}

			if tmp[currPos] > 9 {
				if nextPos > len(tmp)-1 {
					tmp = append(tmp, tmp[currPos]/10)
				} else {
					tmp[nextPos] += tmp[currPos] / 10
				}
				tmp[currPos] %= 10
			}
		}
	}

	result := make([]byte, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		result[len(tmp)-1-i] = tmp[i] + 48
	}

	return string(result)
}
