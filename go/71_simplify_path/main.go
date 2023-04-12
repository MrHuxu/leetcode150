package main

import "strings"

// code
func simplifyPath(path string) string {
	var stack []string
	for _, dir := range strings.Split(path, "/") {
		if dir == "" || dir == "." {
			continue
		}
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, dir)
	}
	return "/" + strings.Join(stack, "/")
}
