// MultiReturnValueFunctions
package main

import (
	"fmt"
)

func multiReturnValueFunction(i, j, k int) (int, int) {
	return i + j + k, i * j * k
}

func main() {
	p, q := multiReturnValueFunction(1, 2, 3)
	fmt.Println(p, q)
}
