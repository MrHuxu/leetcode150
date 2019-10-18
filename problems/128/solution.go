package leetcode150

// code
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 1

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
			if tail-head+1 > result {
				result = tail - head + 1
			}
			mapHeadToTail[head] = tail
			mapTailToHead[tail] = head

		case ok1:
			if tail-num+1 > result {
				result = tail - num + 1
			}
			mapHeadToTail[num] = tail
			mapTailToHead[tail] = num

		case ok2:
			if num-head+1 > result {
				result = num - head + 1
			}
			mapHeadToTail[head] = num
			mapTailToHead[num] = head

		default:
			mapHeadToTail[num] = num
			mapTailToHead[num] = num
		}
	}

	return result
}
