package main

// code
func generateParenthesis(n int) []string {
	return helper(1, 0, n, "(")
}

func helper(numLeft, matchedLeft, n int, pre string) []string {
	if numLeft < n {
		tmp := helper(numLeft+1, matchedLeft, n, pre+"(")
		if matchedLeft < numLeft {
			tmp = append(tmp, helper(numLeft, matchedLeft+1, n, pre+")")...)
		}

		return tmp
	}

	if matchedLeft < numLeft {
		return helper(numLeft, matchedLeft+1, n, pre+")")
	}
	return []string{pre}
}
