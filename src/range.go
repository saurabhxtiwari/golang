// range is used to iterate over elements of collection - array, slice, map, string etc
package main

import (
	"fmt"
)

func main() {

	array := [3]string{"i", "love", "go"}

	for i, val := range array {
		fmt.Println(i)
		fmt.Println(val)

		if val == "love" {
			fmt.Println("found")
		}
	}

	m := map[string]int{"foo": 1, "bar": 2}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

}
