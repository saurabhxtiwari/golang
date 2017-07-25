// Panic
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Before error")
	panic("Unknown error")
}
