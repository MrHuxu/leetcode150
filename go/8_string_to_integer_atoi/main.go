package main

// code
func myAtoi(str string) int {
	num := 0

	sym := 1
	started := false
	for i := 0; i < len(str); i++ {
		b := str[i]

		if started {
			if b >= 48 && b <= 57 {
				num = num*10 + (int(b)-48)*sym

				if num > 2147483647 {
					return 2147483647
				} else if num < -2147483648 {
					return -2147483648
				}
			} else {
				break
			}
		} else {
			if b == 32 {
				continue
			} else if b == 43 {
				started = true
			} else if b == 45 {
				started = true
				sym = -1
			} else if b >= 48 && b <= 57 {
				started = true
				num = sym * (int(b) - 48)
			} else {
				break
			}
		}
	}

	return num
}
