package main

import "fmt"

func main() {
	a := 1

	func() {
		defer func() { a = 2 }()
		defer func() { a = 3 }()
	}()

	fmt.Printf("a=%d\n", a)
	fmt.Printf("value()=%d\n", value())
}

// defer can change a named return value
func value() (a int) {
	a = 1

	defer func() { a = 2 }()
	defer func() { a = 3 }()

	return
}
