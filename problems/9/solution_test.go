package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isPalindrome(121))
	assert.Equal(false, isPalindrome(-121))
	assert.Equal(false, isPalindrome(10))
	assert.Equal(true, isPalindrome(0))
}
