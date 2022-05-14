package main

import (
	"fmt"
)

func main() {
	var stringTree tree[string]
	stringTree.value = "root"
	fmt.Printf("tree: %v\n", stringTree)
}

type tree[T any] struct {
	left, right *tree[T]
	value       T
}

func (t *tree[T]) lookup(x T) *tree[T] { return nil /* FIXME */ }
