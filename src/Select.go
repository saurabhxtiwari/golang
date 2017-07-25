// Select
package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "MessageOne"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "MessageTwo"
	}()

	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			{
				fmt.Println(x)
			}
		case y := <-ch2:
			{
				fmt.Println(y)
			}

		case <-time.After(time.Second):
			{
				fmt.Println("timeout 2")
				return
			}
		}
	}
}
