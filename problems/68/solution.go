package leetcode150

// code
func fullJustify(words []string, maxWidth int) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		var totalLen int

		j := i
		for ; j < len(words); j++ {
			totalLen += len(words[j])

			if totalLen+(j-i) == maxWidth {
				break
			} else if totalLen+(j-i) > maxWidth {
				totalLen -= len(words[j])
				j--
				break
			}
		}

		result = append(result, mergeWords(words, i, j, totalLen, maxWidth))
		i = j
	}

	return result
}

func mergeWords(words []string, i, j, letterLen, width int) string {
	var result string

	switch {
	case i == j:
		result = words[i]

	case j == len(words):
		result = words[i]
		i++
		for ; i < j; i++ {
			result += " " + words[i]
		}

	default:
		gap := (width - letterLen) / (j - i)
		mod := (width - letterLen) % (j - i)
		println(i, j, width, gap, mod)

		result = words[i]
		i++
		for ; i <= j; i++ {
			for k := 0; k < gap; k++ {
				result += " "
			}
			if mod > 0 {
				result += " "
				mod--
			}

			result += words[i]
		}
	}

	for len(result) < width {
		result += " "
	}
	return result
}
