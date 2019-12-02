package main

// code
import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	eval := func(sym string) {
		num2 := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		num1 := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		switch sym {
		case "+":
			stack = append(stack, num1+num2)
		case "-":
			stack = append(stack, num1-num2)
		case "*":
			stack = append(stack, num1*num2)
		case "/":
			stack = append(stack, num1/num2)

		}
	}

	for _, token := range tokens {
		if isSymbol(token) {
			eval(token)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func isSymbol(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/"
}
