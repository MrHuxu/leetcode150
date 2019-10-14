package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	rotate(matrix)
	assert.Equal([][]int{
		[]int{7, 4, 1},
		[]int{8, 5, 2},
		[]int{9, 6, 3},
	}, matrix)

	matrix = [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(matrix)
	assert.Equal([][]int{
		[]int{15, 13, 2, 5},
		[]int{14, 3, 4, 1},
		[]int{12, 6, 8, 9},
		[]int{16, 7, 10, 11},
	}, matrix)
}
