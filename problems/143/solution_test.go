package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	list := BuildList([]int{1, 2, 3, 4})
	reorderList(list)
	assert.Equal(BuildList([]int{1, 4, 2, 3}), list)

	list2 := BuildList([]int{1, 2, 3, 4, 5})
	reorderList(list2)
	assert.Equal(BuildList([]int{1, 5, 2, 4, 3}), list2)
}
