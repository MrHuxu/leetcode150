package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		swapPairs(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		),
	)
}
