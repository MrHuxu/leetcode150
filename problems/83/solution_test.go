package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2}),
		deleteDuplicates(BuildList([]int{1, 1, 2})),
	)

	assert.Equal(
		BuildList([]int{1, 2, 3}),
		deleteDuplicates(BuildList([]int{1, 1, 2, 3, 3})),
	)
}
