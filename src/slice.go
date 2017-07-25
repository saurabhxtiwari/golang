package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := make([]string, 10)
	fmt.Println(s)

	c := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		c[i] = strconv.Itoa(i)
	}

	fmt.Println(c)

	c = append(c, "10", "11")

	fmt.Println(c)

	p := make([][]string, 3)

	fmt.Println(p)

	for i := 0; i < 3; i++ {
		p[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			p[i][j] = strconv.Itoa(i + j)
		}
	}

	fmt.Println(p)

	fmt.Println(c[2:5])

}
