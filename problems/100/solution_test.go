package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		false, isSameTree(
			BuildTree([]interface{}{1, 2, 1}),
			BuildTree([]interface{}{1, 1, 2}),
		),
	)
}
