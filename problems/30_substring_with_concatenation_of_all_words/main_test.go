package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSubstring(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 9}, findSubstring(
		"barfoothefoobarman", []string{"foo", "bar"},
	))
	assert.Equal([]int(nil), findSubstring(
		"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"},
	))
}
