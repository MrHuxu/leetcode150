package main

// code
func generateParenthesis(n int) []string {
	return generateParenthesisHelper(1, 0, n, "(")
}

func generateParenthesisHelper(numLeft, matchedLeft, n int, pre string) []string {
	if numLeft < n {
		tmp := generateParenthesisHelper(numLeft+1, matchedLeft, n, pre+"(")
		if matchedLeft < numLeft {
			tmp = append(tmp, generateParenthesisHelper(numLeft, matchedLeft+1, n, pre+")")...)
		}

		return tmp
	}

	if matchedLeft < numLeft {
		return generateParenthesisHelper(numLeft, matchedLeft+1, n, pre+")")
	}
	return []string{pre}
}
