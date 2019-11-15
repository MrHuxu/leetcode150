package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		11, minimumTotal([][]int{
			[]int{2},
			[]int{3, 4},
			[]int{6, 5, 7},
			[]int{4, 1, 8, 3},
		}),
	)
}
