package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_exist(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCCED"))

	assert.Equal(true, exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "SEE"))

	assert.Equal(false, exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}
