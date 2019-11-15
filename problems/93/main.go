package main

import "strings"

// code
func restoreIpAddresses(s string) []string {
	return recursion(nil, 0, s)
}

func recursion(pre []string, index int, s string) []string {
	if index == len(s) && len(pre) == 4 {
		return []string{strings.Join(pre, ".")}
	}

	if len(s)-index < 4-len(pre) || len(s)-index > (4-len(pre))*3 {
		return nil
	}

	result := recursion(
		append([]string{}, append(pre, string([]byte{s[index]}))...), index+1, s,
	)

	if index+1 < len(s) && s[index] != '0' {
		result = append(
			result,
			recursion(
				append([]string{}, append(pre, string([]byte{s[index], s[index+1]}))...), index+2, s,
			)...,
		)
	}

	if index+2 < len(s) &&
		(s[index] == '1' || (s[index] == '2' && (s[index+1] < '5' || (s[index+1] == '5' && s[index+2] <= '5')))) {
		result = append(
			result,
			recursion(
				append([]string{}, append(pre, string([]byte{s[index], s[index+1], s[index+2]}))...), index+3, s,
			)...,
		)
	}

	return result
}
