package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10

	switch i {
	case 1:
		fmt.Println("i is one")
		break
	case 10:
		fmt.Println("i is ten")
		break
	default:
		fmt.Println("i is not specified")
	}

	p := time.Now().Weekday()

	switch p {
	case time.Monday:
		fmt.Println("Its Monday")
	default:
		fmt.Println("It seems to be some other day::")
	}
}
