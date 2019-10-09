package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"}, letterCombinations("23"),
	)

	assert.Equal(
		[]string{"adg", "bdg", "cdg", "aeg", "beg", "ceg", "afg", "bfg", "cfg", "adh", "bdh", "cdh", "aeh", "beh", "ceh", "afh", "bfh", "cfh", "adi", "bdi", "cdi", "aei", "bei", "cei", "afi", "bfi", "cfi"}, letterCombinations("234"),
	)
}
