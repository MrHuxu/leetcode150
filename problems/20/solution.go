package main

// code
var mapRightToLeft = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}

		if left, ok := mapRightToLeft[s[i]]; ok {
			if stack[len(stack)-1] == left {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
