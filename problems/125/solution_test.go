package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isPalindrome("A man, a plan, a canal: Panama"))
	assert.Equal(false, isPalindrome("race a car"))
	assert.Equal(false, isPalindrome("0P"))
}
