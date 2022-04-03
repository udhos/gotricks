package main

import "fmt"

func main() {
	fmt.Printf("use %%t for bool:       true=%t\n", true)
	fmt.Printf("use %%b for base 2:     15=%b\n", 15)
	fmt.Printf("use %%v for default:    []int{1}=%v\n", []int{1})
	fmt.Printf("use %%#v for go-syntax: []int{1}=%#v\n", []int{1})
	fmt.Printf("use %%T for type:       []int{1}=%T\n", []int{1})
	fmt.Println()
	fmt.Printf("use %%* before verb for width: %%*s with 5 and 'abc': [%*s]\n", 5, "abc")
	fmt.Println()
	fmt.Printf("use %%[i] before verb to pick an argment: value=%[1]v (%%[1]v) type=%[1]T (%%[1]T)\n", 10)
}
