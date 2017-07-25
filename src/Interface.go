// Interface
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectange struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectange) area() float64 {
	return r.length * r.width
}

func (r rectange) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func print(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	c := circle{radius: 3.24}
	r := rectange{length: 1.25, width: 1.25}

	print(c)
	print(r)
}
