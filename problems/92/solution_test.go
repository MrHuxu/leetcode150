package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 4, 3, 2, 5}),
		reverseBetween(
			BuildList([]int{1, 2, 3, 4, 5}), 2, 4,
		),
	)
}
