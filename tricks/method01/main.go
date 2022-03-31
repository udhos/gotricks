package main

import "fmt"

type number float64

func (n number) add(m number) number {
	return n + m
}

func main() {

	n := number(1)

	addToN := n.add // the selector n.add (var.method) issues a "method value"

	fmt.Printf("addToN is a 'method value' from n.add: addtoN(2) => %g\n", addToN(2))

	meAdd := number.add // the selector number.add (type.method) issues a "method expression"

	fmt.Printf("meAdd is a 'method expression' from number.add: meAdd(2,3) => %g\n", meAdd(2, 3))
}
