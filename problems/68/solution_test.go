package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}, fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))

	assert.Equal([]string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}, fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))

	assert.Equal([]string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}, fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}
