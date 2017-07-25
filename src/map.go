package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println(m)

	fmt.Println(m["k2"])

	delete(m, "k2")
	_, p := m["k2"]
	fmt.Println(p)
}
