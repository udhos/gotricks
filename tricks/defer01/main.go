package main

import "fmt"

func main() {
	fmt.Println(returnOne())
}

// return 1 without any return
func returnOne() (one int) {
	defer func() {
		recover()
		one = 1
	}()
	panic("ugh")
}
