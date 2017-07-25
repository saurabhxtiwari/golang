package main

import (
	"fmt"
)

func main() {

	var i [5]int

	fmt.Println(i)

	i[4] = 100

	fmt.Println(i)

	j := [5]int{1, 2, 3, 4, 5}

	fmt.Println(j)

	var k [4][4]int

	for p := 0; p < 4; p++ {
		for q := 0; q < 4; q++ {
			fmt.Println(k[p][q])
		}
	}
}
