package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numTrees(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, numTrees(3))
	assert.Equal(0, numTrees(0))
	assert.Equal(1, numTrees(1))
	assert.Equal(2, numTrees(2))
}
