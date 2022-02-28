package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
  assert := assert.New(t)

  assert.Equal(5, lengthOfLastWord("Hello World"))
}
