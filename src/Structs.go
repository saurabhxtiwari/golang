// Structs
package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
}

func main() {
	p := Person{age: 10, name: "chuju"}

	fmt.Println(p)
	fmt.Println(p.age)

	i := &p
	fmt.Println(i.age)
}
