package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	assert := assert.New(t)

	// assert.Equal(
	// 	"PAHNAPLSIIGYIR",
	// 	convert("PAYPALISHIRING", 3),
	// )
	// assert.Equal(
	// 	"PINALSIGYAHRPI",
	// 	convert("PAYPALISHIRING", 4),
	// )
	assert.Equal(
		"AB",
		convert("AB", 1),
	)
}
