package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_minDepth(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, minDepth(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
	assert.Equal(2, minDepth(BuildTree([]interface{}{1, 2})))
}
