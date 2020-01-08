package main

// code
func candy1(ratings []int) int {
	var result int

	var downward, curr int
	for i, rating := range ratings {
		if i == 0 {
			curr = 1
		} else {
			switch {
			case rating > ratings[i-1]:
				downward = 0
				curr++

			case rating == ratings[i-1]:
				downward = 0
				curr = 1

			case rating < ratings[i-1]:
				downward++
				if curr == 1 {
					result += downward
				} else {
					downward += curr - 1 - 1
				}
				curr = 1
			}
		}
		result += curr
	}

	return result
}

func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		switch {
		case ratings[i] > ratings[i-1]:
			candies[i] = candies[i-1] + 1

		case ratings[i] < ratings[i-1]:
			candies[i] = 1

		default:
			candies[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		switch {
		case ratings[i] > ratings[i+1]:
			if candies[i] < candies[i+1]+1 {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	var result int
	for _, c := range candies {
		result += c
	}
	return result
}
