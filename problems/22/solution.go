package leetcode150

// code
func generateParenthesis(n int) []string {
	return gen(1, 0, n, "(")
}

func gen(numLeft, matchedLeft, total int, pre string) []string {
	if numLeft < total {
		tmp := gen(numLeft+1, matchedLeft, total, pre+"(")
		if matchedLeft < numLeft {
			tmp = append(tmp, gen(numLeft, matchedLeft+1, total, pre+")")...)
		}

		return tmp
	}

	if matchedLeft < numLeft {
		return gen(numLeft, matchedLeft+1, total, pre+")")
	}
	return []string{pre}
}
