package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_grayCode(t *testing.T) {
  assert := assert.New(t)

  assert.Equal([]int{0, 1, 3, 2}, grayCode(2))
}
