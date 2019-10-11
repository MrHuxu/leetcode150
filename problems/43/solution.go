package leetcode150

import "fmt"

// code
func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var lines []string

	for i := len(num2) - 1; i >= 0; i-- {
		lines = append(lines, singleMulti(num1, num2[i], len(num2)-1-i))
	}
	fmt.Println(lines)

	for len(lines) > 1 {
		var next []string
		for i := 0; i < len(lines); i += 2 {
			if i+1 < len(lines) {
				next = append(next, add(lines[i], lines[i+1]))
			} else {
				next = append(next, lines[i])
			}
		}
		lines = next
	}

	return lines[0]
}

func singleMulti(num1 string, num2 byte, pow int) string {
	var bytes []byte
	for i := 0; i < pow; i++ {
		bytes = append(bytes, '0')
	}

	var plus byte
	for i := len(num1) - 1; i >= 0; i-- {
		product := (num1[i] - 48) * (num2 - 48)

		product += plus
		plus = product / 10
		product = product%10 + 48

		bytes = append([]byte{product}, bytes...)
	}
	if plus > 0 {
		bytes = append([]byte{plus + 48}, bytes...)
	}

	return string(bytes)
}

func add(num1, num2 string) string {
	var bytes []byte

	var plus byte
	i := len(num1) - 1
	j := len(num2) - 1
	for i >= 0 || j >= 0 {
		var sum byte
		if i >= 0 {
			sum += num1[i] - 48
			i--
		}
		if j >= 0 {
			sum += num2[j] - 48
			j--
		}

		sum += plus
		plus = sum / 10
		sum = sum%10 + 48

		bytes = append([]byte{sum}, bytes...)
	}
	if plus == 1 {
		bytes = append([]byte{'1'}, bytes...)
	}

	return string(bytes)
}
