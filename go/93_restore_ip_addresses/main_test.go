package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreIpAddresses(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]string{"255.255.11.135", "255.255.111.35"},
		restoreIpAddresses("25525511135"),
	)
}
