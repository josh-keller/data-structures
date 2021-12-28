package main

import (
	"fmt"

	// stack "github.com/wboard82/data-structures/llstack"
	stack "github.com/wboard82/data-structures/sliceStack"
)

func main() {
	s := stack.NewStack()

	s.Push(1)
	fmt.Println(s)
}
