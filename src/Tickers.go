// Tickers
package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.NewTicker(time.Second * 2)

	go func() {
		for {
			<-tick.C
			fmt.Println("Event tiggered")
		}
	}()

	time.Sleep(time.Second * 10)
	tick.Stop()
}
