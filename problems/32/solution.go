package leetcode150

import "fmt"

// code
func longestValidParentheses(s string) int {
	stack := []byte{}
	var currLen byte

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			currLen = 0
			continue
		}

		switch s[i] {
		case '(':
			stack = append(stack, currLen+50, '(')
			currLen = 0

		case ')':
			var foundLeft bool
			var backLen byte
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == ')' {
					break
				} else if stack[j] == '(' {
					stack = stack[0:j]
					stack = append(stack, backLen+2+50)
					foundLeft = true
					break
				} else {
					backLen += stack[j] - 50
				}
			}
			if !foundLeft {
				stack = append(stack, ')')
			}
		}
	}
	stack = append(stack, currLen)
	fmt.Println(stack)

	var tmp, result int
	for _, item := range stack {
		if item == 0 {
			continue
		}

		switch item {
		case '(', ')':
			if tmp > result {
				result = tmp
			}
			tmp = 0

		default:
			tmp += int(item - 50)
		}
	}
	if tmp > result {
		result = tmp
	}

	return result
}
