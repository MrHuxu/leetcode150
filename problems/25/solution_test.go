package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{2, 1, 4, 3, 5}),
		reverseKGroup(
			BuildList([]int{1, 2, 3, 4, 5}), 2,
		),
	)

	assert.Equal(
		BuildList([]int{3, 2, 1, 4, 5}),
		reverseKGroup(
			BuildList([]int{1, 2, 3, 4, 5}), 3,
		),
	)
}
