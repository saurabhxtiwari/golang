// JSON
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int
	Mame string
}

type PersonAgain struct {
	Age  int    `json:"personAge"`
	Name string `json:"personName"`
}

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	p1 := Person{Age: 10, Mame: "saurabh"}
	p2 := PersonAgain{Age: 20, Name: "shobhit"}

	b, error := json.Marshal(&p1)
	if error != nil {
		panic(error)
	}

	c, er := json.Marshal(&p2)
	if er != nil {
		panic(er)
	}

	fmt.Println(string(b))
	fmt.Println(string(c))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	d := Person{}
	json.Unmarshal(b, &d)
	fmt.Println(d.Age)
}
