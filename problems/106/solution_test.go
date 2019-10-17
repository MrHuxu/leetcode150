package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
		buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}),
	)
}
