package main

// code
var mapNumToSym = map[int]map[int]string{
	0: map[int]string{
		1:  "I",
		5:  "V",
		10: "X",
	},
	1: map[int]string{
		1:  "X",
		5:  "L",
		10: "C",
	},
	2: map[int]string{
		1:  "C",
		5:  "D",
		10: "M",
	},
	3: map[int]string{
		1: "M",
	},
}

func intToRoman(num int) string {
	var result string

	var level int
	for num > 0 {
		tmp := num % 10

		switch {
		case tmp > 0 && tmp <= 3:
			for tmp > 0 {
				result = mapNumToSym[level][1] + result
				tmp--
			}

		case tmp == 4:
			result = mapNumToSym[level][1] + mapNumToSym[level][5] + result

		case tmp > 4 && tmp <= 8:
			for tmp > 5 {
				result = mapNumToSym[level][1] + result
				tmp--
			}
			result = mapNumToSym[level][5] + result

		case tmp == 9:
			result = mapNumToSym[level][1] + mapNumToSym[level][10] + result
		}

		level++
		num /= 10
	}

	return result
}
