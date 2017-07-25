// GoRoutines
package main

import (
	"fmt"
)

func printAsync() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go printAsync()

	go printAsync()

	fmt.Println("Done")
}
