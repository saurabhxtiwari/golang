// varArgsFunctions
package main

import (
	"fmt"
)

func varArgsFunc(numbers ...int) (int, int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return len(numbers), sum
}

func main() {
	fmt.Println(varArgsFunc(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
