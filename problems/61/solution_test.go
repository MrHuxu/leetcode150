package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		rotateRight(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			}, 2,
		),
	)

	assert.Equal(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		rotateRight(
			&ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			}, 4,
		),
	)
}
