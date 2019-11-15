package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	input := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(input)
	assert.Equal([][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}, input)

	input2 := [][]byte{
		[]byte{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		[]byte{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'},
		[]byte{'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'},
		[]byte{'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
		[]byte{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'},
		[]byte{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
	}
	solve(input2)
	assert.Equal([][]byte{
		[]byte{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		[]byte{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		[]byte{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		[]byte{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		[]byte{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
	}, input2)
}
