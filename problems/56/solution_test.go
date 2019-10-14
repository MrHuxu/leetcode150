package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 6},
		[]int{8, 10},
		[]int{15, 18},
	}, merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}))
}
