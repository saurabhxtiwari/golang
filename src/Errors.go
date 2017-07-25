// Errors
package main

import (
	"fmt"
)

type MyError struct {
	code    int
	message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}

type Person struct {
	age  float64
	name string
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		e := MyError{code: 1, message: "divide by 0 error"}
		return 0.0, e
	} else {
		return a / b, nil
	}
}

func main() {
	p := Person{age: 10, name: "test"}
	q, e := divide(p.age, 5)
	fmt.Println(q)
	fmt.Println(e)

	q, e = divide(p.age, 0)
	fmt.Println(q)
	fmt.Println(e)
}
