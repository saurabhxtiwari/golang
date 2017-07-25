// MethodStructs
package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func main() {
	r := rectangle{length: 10, width: 10}
	fmt.Println(r.area())
}
