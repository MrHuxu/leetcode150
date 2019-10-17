package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		true, isSymmetric(BuildTree([]interface{}{1, 2, 2, 3, 4, 4, 3})),
	)
	assert.Equal(
		false, isSymmetric(BuildTree([]interface{}{1, 2, 2, nil, 3, nil, 3})),
	)
}
