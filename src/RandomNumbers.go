// RandomNumbers
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	r1 := rand.Intn(100)

	r2 := rand.Intn(100)
	fmt.Println(r1)
	fmt.Println(r2)

	seed := rand.NewSource(time.Now().UnixNano())
	r3 := rand.New(seed)
	fmt.Println(r3.Intn(100))

}
