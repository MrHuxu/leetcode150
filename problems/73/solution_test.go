package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	setZeroes(matrix)
	assert.Equal([][]int{
		[]int{1, 0, 1},
		[]int{0, 0, 0},
		[]int{1, 0, 1},
	}, matrix)

	matrix2 := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	setZeroes(matrix2)
	assert.Equal([][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 4, 5, 0},
		[]int{0, 3, 1, 0},
	}, matrix2)
}
