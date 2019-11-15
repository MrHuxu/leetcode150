package main

// code
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int

	var overlapLeft, overlapRight int
	var hasOverlap, overlapUsed bool
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			result = append(result, interval)
			continue
		}

		if interval[0] > newInterval[1] {
			if !overlapUsed {
				if !hasOverlap {
					result = append(result, newInterval)
				} else {
					result = append(result, []int{overlapLeft, overlapRight})
				}
				overlapUsed = true
			}

			result = append(result, interval)
			continue
		}

		if !hasOverlap {
			overlapLeft = min(interval[0], newInterval[0])
			overlapRight = max(interval[1], newInterval[1])
			hasOverlap = true
		} else {
			overlapLeft = min(interval[0], overlapLeft)
			overlapRight = max(interval[1], overlapRight)
		}
	}
	if !overlapUsed {
		if !hasOverlap {
			result = append(result, newInterval)
		} else {
			result = append(result, []int{overlapLeft, overlapRight})
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
