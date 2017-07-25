// Channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	chBuffered := make(chan string, 2)

	go func() {
		ch <- "Message"
	}()

	go func() {
		fmt.Println("Waiting....")
		time.Sleep(time.Second)
		chBuffered <- "Buffered Message One"
		chBuffered <- "Buffered Message Two"
		chBuffered <- "Buffered Message Three"
	}()

	x := <-ch
	fmt.Println(x)

	fmt.Println(<-chBuffered)
	fmt.Println(<-chBuffered)
	fmt.Println(<-chBuffered)

}
