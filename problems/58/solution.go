package main

// code
import "strings"

func lengthOfLastWord(word string) int {
	arr := strings.Split(strings.Trim(word, " "), " ")

	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}
