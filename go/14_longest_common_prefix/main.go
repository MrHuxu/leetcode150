package main

import "strings"

// code
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	ret := strs[0]
	for _, str := range strs {
		for !strings.HasPrefix(str, ret) {
			ret = ret[:len(ret)-1]
		}
	}
	return ret
}
