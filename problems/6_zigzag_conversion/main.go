package main

// code
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := make(map[int][]byte)
	vertical := true
	line := 0
	for i := 0; i < len(s); i++ {
		m[line] = append(m[line], s[i])
		if vertical {
			line++
		} else {
			line--
		}

		if line == numRows-1 {
			vertical = false
		} else if line == 0 {
			vertical = true
		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		result += string(m[i])
	}
	return result
}
