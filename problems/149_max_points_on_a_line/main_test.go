package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, maxPoints([][]int{
		[]int{1, 1},
		[]int{2, 2},
		[]int{3, 3},
	}))

	assert.Equal(4, maxPoints([][]int{
		[]int{1, 1},
		[]int{3, 2},
		[]int{5, 3},
		[]int{4, 1},
		[]int{2, 3},
		[]int{1, 4},
	}))

	assert.Equal(3, maxPoints([][]int{
		[]int{0, 0},
		[]int{1, 1},
		[]int{0, 0},
	}))

	assert.Equal(3, maxPoints([][]int{
		[]int{4, 0},
		[]int{4, -1},
		[]int{4, 5},
	}))

	assert.Equal(4, maxPoints([][]int{
		[]int{3, 10},
		[]int{0, 2},
		[]int{0, 2},
		[]int{3, 10},
	}))
}
