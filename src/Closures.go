// Closures
package main

import (
	"fmt"
)

func nextInt() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	i := nextInt()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}
