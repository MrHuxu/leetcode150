package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
		addTwoNumbers(
			&ListNode{Val: 5},
			&ListNode{Val: 5},
		),
	)
	assert.Equal(
		&ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
		addTwoNumbers(
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		),
	)
}
