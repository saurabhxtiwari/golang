// ClosingChannels
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	ch := make(chan string)
	doneCh := make(chan bool)

	go func() {
		for {
			j, hasMoreData := <-ch
			if hasMoreData {
				fmt.Println("Job received on channel ", j)
			} else {
				fmt.Println("All jobs completed")
				doneCh <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 3)
		ch <- ("Job " + strconv.Itoa(i))
	}

	close(ch)

	<-doneCh
}
