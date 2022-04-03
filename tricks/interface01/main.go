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
	//   the underlying object has a type and a value
	//     in case of a pointer type, it can be nil
	//     HOWEVER a nil interface is different than an interface holding a nil pointer!
	//     example:
	//       type:  *bytes.Buffer
	//       value: nil

	fmt.Println()
	fmt.Println("interface holding nil pointer to *bytes.Buffer:")
	fmt.Println()

	var buf *bytes.Buffer
	iface1 = buf

	fmt.Printf("iface1 is NOT nil:              type=%[1]T value=%[1]v\n", iface1)
	fmt.Printf("iface1 == nil:                  %t\n", iface1 == nil)
	fmt.Printf("iface1 == (*bytes.Buffer)(nil): %t\n", iface1 == (*bytes.Buffer)(nil))

	fmt.Println()
	fmt.Println("let's take a look at an interface holding a non-pointer object:")
	fmt.Println()

	var iface2 interface{} = [...]int{1, 2}
	fmt.Printf("iface2: type=%[1]T value=%[1]v\n", iface2)
}
