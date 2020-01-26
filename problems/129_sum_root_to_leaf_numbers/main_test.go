package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(25, sumNumbers(BuildTree([]interface{}{1, 2, 3})))
	assert.Equal(1026, sumNumbers(BuildTree([]interface{}{4, 9, 0, 5, 1})))
}
