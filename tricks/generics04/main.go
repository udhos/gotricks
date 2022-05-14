package main

import (
	"fmt"
)

func main() {
	fmt.Printf("work1: %t\n", work1([]string{}, "a"))
	fmt.Printf("work2: %t\n", work2([]string{}, "a"))
	fmt.Printf("work3: %t\n", work3([]string{}, "a"))
}

// Interfaces used as constraints may be given names (such as Ordered),
// or they may be literal interfaces inlined in a type parameter list.
// For example:
// [S interface{~[]E}, E interface{}]
func work1[S interface{ ~[]E }, E interface{}](s S, e E) bool {
	return false
}

// Because this is a common case, the enclosing interface{} may be omitted
// for interfaces in constraint position, and we can simply write:
// [S ~[]E, E interface{}]
func work2[S ~[]E, E interface{}](s S, e E) bool {
	return false
}

// Because the empty interface is common in type parameter lists, and in
// ordinary Go code for that matter, Go 1.18 introduces a new predeclared
// identifier any as an alias for the empty interface type.
// With that, we arrive at this idiomatic code:
// [S ~[]E, E any]
func work3[S ~[]E, E any](s S, e E) bool {
	return false
}
