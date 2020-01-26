package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_minDepth(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, minDepth(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
	assert.Equal(2, minDepth(BuildTree([]interface{}{1, 2})))
}
