// Timer
package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTimer(time.Second * 3)
	t2 := time.NewTicker(time.Second * 4)

	for {
		x := <-t1.C
		fmt.Println(x)
		break
		/*if x {
			fmt.Println("Timer completed")
			return
		}*/
	}

	go func() {
		fmt.Println("Waiting for t2 to expire")
		<-t2.C
	}()

	time.Sleep(time.Second * 2)
	t2.Stop()

	fmt.Println("t2 interrupted")

}
