package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, romanToInt("III"))
	assert.Equal(4, romanToInt("IV"))
	assert.Equal(9, romanToInt("IX"))
	assert.Equal(58, romanToInt("LVIII"))
	assert.Equal(1994, romanToInt("MCMXCIV"))
}
