package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2, 2, 4, 3, 5}),
		partition(BuildList([]int{1, 4, 3, 2, 5, 2}), 3),
	)
}
