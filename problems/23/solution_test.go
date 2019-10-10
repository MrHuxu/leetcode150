package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
			},
		},

		mergeKLists(
			[]*ListNode{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		),
	)
}
