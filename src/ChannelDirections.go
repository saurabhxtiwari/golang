// ChannelDirections
package main

import (
	"fmt"
	"time"
)

func receiverChannel(chreceiver chan<- string, msg string) {
	chreceiver <- msg
}

func senderChannel(chsender <-chan string, chreceiver chan<- string) {
	x := <-chsender
	time.Sleep(time.Second)
	chreceiver <- x
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go receiverChannel(ch1, "Message")
	go senderChannel(ch1, ch2)

	fmt.Println(<-ch2)
}
