package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[][]int{
			[]int{5, 4, 11, 2},
			[]int{5, 8, 4, 5},
		},
		pathSum(BuildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), 22),
	)
}
