package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Printf("instantiate as gmin[int]:     gmin(3,4): %d\n", gmin(3,4))
	fmt.Printf(`instantiate as gmin[string]:  gmin("a","b"): %s`+"\n", gmin("a","b"))

	fmin := gmin[float64]

	fmt.Printf("instantiate as gmin[float64]: fmin(1.1,2.2): %g\n", fmin(1.1,2.2))
}

func gmin[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}
