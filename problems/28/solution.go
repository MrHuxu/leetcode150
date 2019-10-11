package leetcode150

// code
func strStr(haystack, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		found := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				found = false
			}
		}

		if found {
			return i
		}
	}
	return -1
}
