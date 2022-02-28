package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{1, 4},
		[]int{2, 3},
		[]int{2, 4},
		[]int{3, 4},
	}, combine(4, 2))

	assert.Equal([][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 5},
		[]int{1, 2, 4, 5},
		[]int{1, 3, 4, 5},
		[]int{2, 3, 4, 5},
	}, combine(5, 4))
}
