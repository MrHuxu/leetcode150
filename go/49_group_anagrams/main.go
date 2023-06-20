package main

// code
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := count(str)
		groups[cnt] = append(groups[cnt], str)
	}

	ret := make([][]string, 0, len(groups))
	for _, group := range groups {
		ret = append(ret, group)
	}
	return ret
}

func count(str string) (ret [26]int) {
	for _, ch := range str {
		ret[ch-'a']++
	}
	return
}
