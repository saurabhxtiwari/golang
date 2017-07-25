// Pointers
package main

import (
	"fmt"
)

func noPointer(i int) {
	i = 0
}

func pointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	noPointer(i)
	fmt.Println(i)

	pointer(&i)
	fmt.Println(i)
}
