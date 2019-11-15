package main

// code
func isScramble(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	l := len(s1)

	if l == 1 {
		return s1 == s2
	}

	if l == 2 {
		return (s1[0] == s2[0] && s1[1] == s2[1]) || (s1[0] == s2[1] && s1[1] == s2[0])
	}

	cnt := make([]int, 26)
	for i := 0; i < l; i++ {
		cnt[s1[i]-'a']++
		cnt[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			return false
		}
	}

	for i := 1; i < l; i++ {
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:l], s2[i:l]) {
			return true
		}

		if isScramble(s1[0:i], s2[l-i:l]) && isScramble(s1[i:l], s2[0:l-i]) {
			return true
		}
	}

	return false
}
