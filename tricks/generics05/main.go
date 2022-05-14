package main

import (
	"fmt"
)

func main() {
	fmt.Print(`
Function argument type inference only works for type parameters that
are used in the function parameters, not for type parameters used only
in function results or only in the function body.

`)
	fmt.Print("function argument type inference:\n\n")
	fmt.Printf("    possible: work1(3): %v\n", work1(3))
	fmt.Printf("NOT possible: work2[int](): %v\n", work2[int]())
	fmt.Printf("NOT possivel: work3[string](): %v\n", work3[string]())
}

type scalar interface {
	int | string
}

func work1[T scalar](x T) bool {
	return true
}

func work2[T scalar]() T {
	return T(2)
}

func work3[T scalar]() bool {
	return true
}
