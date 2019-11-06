package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isScramble("great", "rgeat"))
	assert.Equal(false, isScramble("abcde", "caebd"))
	assert.Equal(true, isScramble("abcd", "dacb"))
	assert.Equal(false, isScramble("bcbbcccbcbcaaacbb", "acbcabbbbacccbbcc"))
}
