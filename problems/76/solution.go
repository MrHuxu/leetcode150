package main

// code
func minWindow(s string, t string) string {
	m := make(map[byte]int)
	for _, ch := range t {
		m[byte(ch)]++
	}
	cnt := len(t)

	var resultL, resultR int
	var found bool
	resultLen := len(s)

	var left, right int
	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			m[s[right]]--

			if m[s[right]] >= 0 {
				cnt--
			}

			if cnt == 0 {
				found = true
			}

			for cnt == 0 {
				if right-left-1 < resultLen {
					resultLen = right - left - 1
					resultL = left
					resultR = right
				}

				if _, ok := m[s[left]]; ok {
					m[s[left]]++

					if m[s[left]] > 0 {
						cnt++
					}
				}

				left++
			}
		}

		right++
	}

	if found {
		return s[resultL : resultR+1]
	}
	return ""
}
