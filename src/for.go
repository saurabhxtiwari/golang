package main

import (
	"fmt"
)

func main() {
	i := 10

	for i > 0 {
		fmt.Println(i)
		i--
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	p := 4
	for {
		if p == 10 {
			break
		}
		fmt.Println(p)
		p++
	}
}
