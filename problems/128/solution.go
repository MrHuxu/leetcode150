package leetcode150

// code
func longestConsecutive(nums []int) int {
	var result int

	mapHeadToTail := make(map[int]int)
	mapTailToHead := make(map[int]int)
	existance := make(map[int]bool)
	for _, num := range nums {
		if existance[num] {
			continue
		} else {
			existance[num] = true
		}

		tail, ok1 := mapHeadToTail[num+1]
		head, ok2 := mapTailToHead[num-1]

		switch {
		case ok1 && ok2:
			result = max(result, tail-head+1)
			mapHeadToTail[head] = tail
			mapTailToHead[tail] = head

		case ok1:
			result = max(result, tail-num+1)
			mapHeadToTail[num] = tail
			mapTailToHead[tail] = num

		case ok2:
			result = max(result, num-head+1)
			mapHeadToTail[head] = num
			mapTailToHead[num] = head

		default:
			result = max(result, 1)
			mapHeadToTail[num] = num
			mapTailToHead[num] = num
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
