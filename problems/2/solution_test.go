package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{0, 1}),
		addTwoNumbers(
			&ListNode{Val: 5},
			&ListNode{Val: 5},
		),
	)
	assert.Equal(
		BuildList([]int{7, 0, 8}),
		addTwoNumbers(
			BuildList([]int{2, 4, 3}),
			BuildList([]int{5, 6, 4}),
		),
	)
}
