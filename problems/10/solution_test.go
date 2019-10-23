package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(false, isMatch("aa", "a"))
	assert.Equal(true, isMatch("aa", "a*"))
	assert.Equal(true, isMatch("ab", ".*"))
	assert.Equal(true, isMatch("aab", "c*a*b"))
	assert.Equal(false, isMatch("mississippi", "mis*is*p*."))
	assert.Equal(true, isMatch("aaa", "ab*ac*a"))
	assert.Equal(false, isMatch("a", ".*..a*"))
	assert.Equal(true, isMatch("a", "ab*a*"))
}
