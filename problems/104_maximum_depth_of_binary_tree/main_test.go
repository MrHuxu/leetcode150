package main

import (
  "testing"

  . "github.com/MrHuxu/types"

  "github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
  assert := assert.New(t)

  assert.Equal(3, maxDepth(BuildTree([]interface{}{3,9,20,nil,nil,15,7})))
}
