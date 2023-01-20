package main

import (
	"fmt"

	"github.com/MrHuxu/leetcode150/website/data"
)

func init() {
	data.Init()
	initServer()
}

func main() {
	Server.Run(fmt.Sprintf(":%d", port))
}
