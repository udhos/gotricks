package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int64{
		"a":1,
		"b":10,
		"c":30,
	}

	fmt.Printf("sumIntsOrFloats(m1): %d\n",sumIntsOrFloats(m1))
	fmt.Printf("sumNumbers(m1): %d\n",sumNumbers(m1))

	m2 := map[string]counterInt{
		"a":1,
		"b":10,
		"c":30,
	}

	fmt.Printf("sumNumbers2(m2): %d\n",sumNumbers2(m2))

	m3 := map[string]counterFloat{
		"a":1.1,
		"b":10,
		"c":30,
	}

	fmt.Printf("sumNumbers2(m3): %g\n",sumNumbers2(m3))
	fmt.Printf("sumNumbers2[string,counterFloat](m3): %g\n",sumNumbers2[string,counterFloat](m3))

	fmt.Printf("numberList[int64]{2}: %v\n", numberList[int64]{2})
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

type number interface {
	int64 | float64
}

type counterInt int64
type counterFloat float64

func sumNumbers[K comparable, V number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

type number2 interface {
	~int64 | ~float64
}

func sumNumbers2[K comparable, V number2](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

type numberList[T number2] []T