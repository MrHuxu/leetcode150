package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}, generate(5))
}
