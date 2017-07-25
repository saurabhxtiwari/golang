package main

import (
	"fmt"
	"math"
)

const (
	s string = "teststring"
)

func main() {

	const n = 500000000

	fmt.Println(s)
	fmt.Println(math.Cos(90))

	const d = 3e20 / n

	fmt.Println(d)
}
