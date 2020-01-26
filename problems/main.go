package main

import (
	"fmt"

	"github.com/MrHuxu/leetcode150/problems/utils"
)

func main() {
	tree := utils.BuildTree([]interface{}{1, 2, 3, 4, 5, 6, 7, nil, nil, 8, 9})
	fmt.Println(tree)
}
