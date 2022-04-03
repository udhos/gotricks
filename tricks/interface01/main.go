package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	fmt.Println("nil interface:")
	fmt.Println()

	var iface1 io.Writer

	fmt.Printf("iface1 is nil: type=%[1]T value=%[1]v\n", iface1)
	fmt.Printf("iface1 == nil: %t\n", iface1 == nil)

	// an interface is either nil or has an underlying object
	// the underlying object can be a nil pointer to a type

	fmt.Println()
	fmt.Println("interface holding nil pointer to *bytes.Buffer:")
	fmt.Println()

	var buf *bytes.Buffer
	iface1 = buf

	fmt.Printf("iface1 is NOT nil:              type=%[1]T value=%[1]v\n", iface1)
	fmt.Printf("iface1 == nil:                  %t\n", iface1 == nil)
	fmt.Printf("iface1 == (*bytes.Buffer)(nil): %t\n", iface1 == (*bytes.Buffer)(nil))
}
