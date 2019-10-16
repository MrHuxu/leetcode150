package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{4, 5, 1, 2, 3}),
		rotateRight(
			BuildList([]int{1, 2, 3, 4, 5}), 2,
		),
	)

	assert.Equal(
		BuildList([]int{2, 0, 1}),
		rotateRight(
			BuildList([]int{0, 1, 2}), 4,
		),
	)
}
