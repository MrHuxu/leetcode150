package main

import (
	"fmt"
)

func init() {
	initData()
	initServer()
}

func main() {
	Server.Run(fmt.Sprintf(":%d", port))
}
