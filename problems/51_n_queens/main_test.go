package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{
		[]string{
			".Q..", // Solution 1
			"...Q",
			"Q...",
			"..Q.",
		},
		[]string{
			"..Q.", // Solution 2
			"Q...",
			"...Q",
			".Q..",
		},
	}, solveNQueens(4))
}
